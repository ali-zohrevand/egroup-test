package model

type GitlabGraphqlResponse struct {
	Projects GitlabProjects `json:"projects"`
}
type GitlabNode struct {
	Description *string `json:"description"`
	ForksCount  int     `json:"forksCount"`
	Name        string  `json:"name"`
}
type GitlabProjects struct {
	Nodes []GitlabNode `json:"nodes"`
}
