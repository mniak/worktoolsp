package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/codefresh-io/go-sdk/pkg/codefresh"
	"github.com/codefresh-io/go-sdk/pkg/rest"
	"github.com/samber/lo"
)

func main() {
	const host = "https://g.codefresh.io"
	const tag = "10.24.1"
	const appName = "hsm-api"

	cf := codefresh.New(&codefresh.ClientOptions{
		Host:  host,
		Token: "67e3205a84626b4ad9fdc529.72fd6d8b4cd15c76d2e5f1e6b8be6e7a",
	})
	api := cf.Rest()
	pipeAPI := api.Pipeline()
	pipelines, err := pipeAPI.List(map[string]string{
		"limit": "40",
		"id":    fmt.Sprintf("%s/deployment-api-", appName),
	})
	if err != nil {
		log.Panic("Failed to list pipelines: ", err)
	}

	pipelines = lo.Filter(pipelines, func(p rest.Pipeline, index int) bool {
		envName, _ := strings.CutPrefix(p.Metadata.Name, "hsm-api/deployment-api-")
		env := Environment(envName)
		// return env.IsValid() && env.IsProduction() && !env.IsMajor()
		return env.Is(India)
	})
	for _, p := range pipelines {
		// buildId, err := pipeAPI.Run(p.Metadata.Name, &rest.RunOptions{
		// 	Variables: map[string]string{
		// 		"GIT_REVISION": tag,
		// 		"TAG":          tag,
		// 	},
		// })
		// if err != nil {
		// 	log.Panic("Failed to list pipelines: ", err)
		// }
		// cfLink := fmt.Sprintf("%s/build/%s", host, buildId)

		cfLink := "didn't start"
		vars := lo.SliceToMap(p.Spec.Variables, func(item CFVariable) (string, string) {
			return item.Key, item.Value
		})
		// pretty.Println(p)
		fmt.Printf("Started %s: %s\n", p.Metadata.Name, cfLink)
		if vars["REGION_PRIMARY_ENABLED"] == "true" {
			argoLink := fmt.Sprintf("https://argocd.pismo.services/applications/%s", vars["REGION_PRIMARY_FULL_NAME"])
			fmt.Println("-> Primary Region ", argoLink)
		}
		if vars["REGION_SECONDARY_ENABLED"] == "true" {
			argoLink := fmt.Sprintf("https://argocd.pismo.services/applications/%s", vars["REGION_SECONDARY_FULL_NAME"])
			fmt.Println("-> Secondary Region ", argoLink)
		}

	}
}

type CFVariable = struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
