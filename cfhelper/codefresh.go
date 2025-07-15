package main

import (
	"fmt"
	"os"

	"github.com/codefresh-io/go-sdk/pkg/codefresh"
)

type CodefreshClient struct {
	codefresh.Codefresh
}

const _CodefreshHost = "https://g.codefresh.io"

func NewCodefreshClient() *CodefreshClient {
	token := os.Getenv("CODEFRESH_TOKEN")

	cf := codefresh.New(&codefresh.ClientOptions{
		Host:  _CodefreshHost,
		Token: token,
	})
	return &CodefreshClient{cf}
}

func (*CodefreshClient) BuildRunURL(runID string) string {
	return fmt.Sprintf("%s/build/%s", _CodefreshHost, runID)
}
