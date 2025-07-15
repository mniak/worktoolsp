package main

import (
	"fmt"
	"os"

	"github.com/codefresh-io/go-sdk/pkg/rest"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func main() {
	cmd := cobra.Command{
		Use:   "cfhelper",
		Short: "Codefresh helper",
	}

	// var flagAppName string
	var flagAppKind string
	var flagGitTag string
	var flagPR string
	deployCmd := cobra.Command{
		Use:  "deploy [options] <App> <Environment>",
		Args: cobra.ExactArgs(2),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if flagAppKind != "api" && flagAppKind != "worker" {
				return fmt.Errorf("Application kind must be 'api' or 'worker'")
			}
			if flagGitTag == "" && flagPR == "" {
				return fmt.Errorf("Either Tag or PR must be specified")
			}
			if flagGitTag != "" && flagPR != "" {
				return fmt.Errorf("Only Tag or PR must be specified")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			app := args[0]
			env, err := ParseEnv(args[1])
			cobra.CheckErr(err)
			if flagPR != "" {
				cobra.CheckErr("PR deployment is not supported yet, please use --tag instead")
			}

			cf := NewCodefreshClient()
			cfAPI := cf.Rest()
			pipelineAPI := cfAPI.Pipeline()

			pipelineName := formatPipelineName(app, flagAppKind, env)
			pipelines, err := pipelineAPI.List(map[string]string{
				"limit": "5",
				"id":    pipelineName,
			})
			handleErr(err, "Failed to list pipelines")

			if len(pipelines) != 1 {
				cobra.CheckErr(fmt.Errorf("found %d pipelines for search %q, expected 1\n", len(pipelines), pipelineName))
			}
			p := pipelines[0]
			result, err := pipelineAPI.Run(p.Metadata.ID, &rest.RunOptions{
				Variables: map[string]string{
					"TAG":          flagGitTag,
					"GIT_REVISION": lo.If(env.Is(Investec), "main").Else(flagGitTag),
				},
			})
			handleErr(err, "Failed to start pipeline")
			fmt.Printf("Started %s: %s\n", p.Metadata.Name, cf.BuildRunURL(result))

			vars := lo.SliceToMap(p.Spec.Variables, func(v CFVariable) (string, string) {
				return v.Key, v.Value
			})
			if vars["REGION_PRIMARY_ENABLED"] == "true" {
				argoLink := fmt.Sprintf("https://argocd.pismo.services/applications/%s", vars["REGION_PRIMARY_FULL_NAME"])
				fmt.Println("-> Primary Region ", argoLink)
			}
			if vars["REGION_SECONDARY_ENABLED"] == "true" {
				argoLink := fmt.Sprintf("https://argocd.pismo.services/applications/%s", vars["REGION_SECONDARY_FULL_NAME"])
				fmt.Println("-> Secondary Region ", argoLink)
			}
		},
	}
	deployCmd.Flags().StringVar(&flagAppKind, "kind", "api", "Application kind (api, worker)")
	deployCmd.Flags().StringVar(&flagGitTag, "tag", "", "Tag to deploy")
	deployCmd.Flags().StringVar(&flagPR, "pr", "", "PR to deploy")

	cmd.AddCommand(&deployCmd)
	cmd.Execute()
}

func formatPipelineName(app, kind string, env Environment) string {
	return fmt.Sprintf("%s-%s/deployment-%s-%s", app, kind,
		kind, env,
	)
}

func handleErr(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s: %s\n", msg, err)
		os.Exit(1)
	}
}

type CFVariable = struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
