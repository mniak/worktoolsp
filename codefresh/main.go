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
	host := "https://g.codefresh.io"
	cf := codefresh.New(&codefresh.ClientOptions{
		Host:  host,
		Token: "67e3205a84626b4ad9fdc529.72fd6d8b4cd15c76d2e5f1e6b8be6e7a",
	})
	api := cf.Rest()
	pipeAPI := api.Pipeline()
	pipelines, err := pipeAPI.List(map[string]string{
		"limit": "40",
		"id":    "hsm-api/deployment-api-",
	})
	if err != nil {
		log.Panic("Failed to list pipelines: ", err)
	}

	const tag = "10.24.1"
	pipelines = lo.Filter(pipelines, func(p rest.Pipeline, index int) bool {
		envName, _ := strings.CutPrefix(p.Metadata.Name, "hsm-api/deployment-api-")
		env := Environment(envName)
		return env.IsValid() && env.IsProduction()
	})
	for _, p := range pipelines {
		log.Print("- ", p.Metadata.Name)
		buildId, err := pipeAPI.Run(p.Metadata.Name, &rest.RunOptions{
			Variables: map[string]string{
				"GIT_REVISION": tag,
				"TAG":          tag,
			},
		})
		if err != nil {
			log.Panic("Failed to list pipelines: ", err)
		}
		link := fmt.Sprintf("%s/build/%s", host, buildId)
		log.Print("Build link: ", link)
	}
}
