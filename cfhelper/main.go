package main

import (
	"fmt"
	"os"

	"github.com/codefresh-io/go-sdk/pkg/rest"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

type Kind string

const (
	API       Kind = "api"
	Worker    Kind = "worker"
	Terraform Kind = "terraform"
)

func main() {
	cmd := cobra.Command{
		Use:   "cfhelper",
		Short: "Codefresh helper",
	}

	// var flagAppName string
	var flagAppKind Kind
	var flagGitTag string
	var flagPR string
	deployCmd := cobra.Command{
		Use:  "deploy [options] <App> <Environment>",
		Args: cobra.ExactArgs(2),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			switch flagAppKind {
			case API, Worker, Terraform:
			case "tf":
				flagAppKind = Terraform
			default:
				return fmt.Errorf("Application kind must be 'api', 'worker' or 'terraform'")
			}
			if flagAppKind != Terraform {
				if flagGitTag == "" && flagPR == "" {
					return fmt.Errorf("Either Tag or PR must be specified")
				}
			} else {
				if flagGitTag != "" {
					return fmt.Errorf("Tag is not supported for Terraform deployments")
				}
				if flagPR != "" {
					return fmt.Errorf("PR deployment is not supported for Terraform")
				}
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
			inputVars := make(map[string]string)
			if flagAppKind != Terraform {
				inputVars["TAG"] = flagGitTag
				inputVars["GIT_REVISION"] = lo.If(env.Is(Investec), "main").Else(flagGitTag)
			}
			result, err := pipelineAPI.Run(p.Metadata.ID, &rest.RunOptions{
				Variables: inputVars,
			})
			handleErr(err, "Failed to start pipeline")
			fmt.Printf("Started %s: %s\n", p.Metadata.Name, cf.BuildRunURL(result))

			defaultValues := lo.SliceToMap(p.Spec.Variables, func(v CFVariable) (string, string) {
				return v.Key, v.Value
			})
			if defaultValues["REGION_PRIMARY_ENABLED"] == "true" {
				argoLink := fmt.Sprintf("https://argocd.pismo.services/applications/%s", defaultValues["REGION_PRIMARY_FULL_NAME"])
				fmt.Println("-> Primary Region ", argoLink)
			}
			if defaultValues["REGION_SECONDARY_ENABLED"] == "true" {
				argoLink := fmt.Sprintf("https://argocd.pismo.services/applications/%s", defaultValues["REGION_SECONDARY_FULL_NAME"])
				fmt.Println("-> Secondary Region ", argoLink)
			}
		},
	}
	deployCmd.Flags().StringVar((*string)(&flagAppKind), "kind", string(API), "Application kind (api, worker, terraform)")
	deployCmd.Flags().StringVar(&flagGitTag, "tag", "", "Tag to deploy")
	deployCmd.Flags().StringVar(&flagPR, "pr", "", "PR to deploy")

	cmd.AddCommand(&deployCmd)
	cmd.Execute()
}

func formatPipelineName(app string, kind Kind, env Environment) string {
	if kind == Terraform {
		return fmt.Sprintf("%s-%s/%s-%s",
			app, API,
			Terraform, env,
		)
	}
	return fmt.Sprintf("%s-%s/deployment-%s-%s",
		app, kind,
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
