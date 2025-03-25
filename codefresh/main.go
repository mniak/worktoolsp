package main

import (
	"log"

	"github.com/codefresh-io/go-sdk/pkg/codefresh"
)

func main() {
	cf := codefresh.New(&codefresh.ClientOptions{
		Host:  "https://g.codefresh.io",
		Token: "67e3205a84626b4ad9fdc529.72fd6d8b4cd15c76d2e5f1e6b8be6e7a",
	})
	api := cf.Rest()
	pipelines, err := api.Pipeline().List(map[string]string{
		"limit": "10",
	})
	if err != nil {
		log.Panic("Failed to list pipelines: ", err)
	}

	log.Print("Pipelines", pipelines)
}
