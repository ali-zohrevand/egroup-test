package repo

import (
	"context"
	"encoding/json"
	"os"
	"strings"

	"github.com/ali-zohrevand/egroup-test.git/model"
	"github.com/machinebox/graphql"
)

func GetLastProjects() (gitlabResponse model.GitlabGraphqlResponse, err error) {
	number := os.Getenv("PROJECTS-NUMBER")
	apiPath := os.Getenv("API-PATH")
	if number == "" {
		number = "5"
	}
	if apiPath == "" {
		apiPath = "https://gitlab.com/api/graphql"
	}
	graphqlClient := graphql.NewClient(apiPath)
	query := `query last_projects($n: Int = **param**) {
			  projects(last:$n) {
				nodes {
				  name
				  description
				  forksCount
				}
			  }
			}`
	query = strings.ReplaceAll(query, "**param**", number)
	graphqlRequest := graphql.NewRequest(query)
	var graphqlResponse interface{}
	if errSendRequest := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); errSendRequest != nil {
		return gitlabResponse, errSendRequest

	}
	res, errUnmarshal := json.Marshal(graphqlResponse)
	if errUnmarshal != nil {
		return gitlabResponse, errUnmarshal
	}
	err = json.Unmarshal(res, &gitlabResponse)
	return gitlabResponse, err

}
