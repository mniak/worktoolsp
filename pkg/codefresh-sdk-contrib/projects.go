package codefreshsdkcontrib

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/codefresh-io/go-sdk/pkg/client"
)

type (
	ProjectAPI interface {
		List(query map[string]string) ([]Project, error)
	}

	project struct {
		client *client.CfClient
	}

	Project struct {
		AccountID       string     `json:"accountId"`
		ProjectName     string     `json:"projectName"`
		UpdatedAt       string     `json:"updatedAt"`
		Metadata        Metadata   `json:"metadata"`
		Image           string     `json:"image"`
		Tags            []string   `json:"tags"`
		Variables       []Variable `json:"variables"`
		PipelinesNumber int64      `json:"pipelinesNumber"`
		ID              string     `json:"id"`
		Favorite        bool       `json:"favorite"`
	}

	Metadata struct {
		CreatedAt string `json:"createdAt"`
	}

	Variable struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
)

func NewProjectAPI(client *client.CfClient) *project {
	return &project{client: client}
}

func (p *project) List(query map[string]string) ([]Project, error) {
	anyQuery := map[string]any{}
	for k, v := range query {
		anyQuery[k] = v
	}

	res, err := p.client.RestAPI(context.TODO(), &client.RequestOptions{
		Method: "GET",
		Path:   "/api/projects",
		Query:  anyQuery,
	})
	if err != nil {
		return nil, fmt.Errorf("failed getting project list: %w", err)
	}

	var result []Project
	return result, json.Unmarshal(res, result)
}
