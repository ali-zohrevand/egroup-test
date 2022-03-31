package DTO

type GitlabLastProjects struct {
	Names         string `json:"names"`
	SumOfAllForks int    `json:"sumOfAllForks"`
}
