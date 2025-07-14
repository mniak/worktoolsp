package main

import (
	"os"

	"github.com/codefresh-io/go-sdk/pkg/codefresh"
)

func NewCodefreshClient() codefresh.Codefresh {
	const host = "https://g.codefresh.io"
	cf := codefresh.New(&codefresh.ClientOptions{
		Host:  host,
		Token: os.Getenv("CODEFRESH_TOKEN"),
	})
	return cf
}
