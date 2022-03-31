package usecase

import (
	"github.com/ali-zohrevand/egroup-test.git/model/DTO"
	"github.com/ali-zohrevand/egroup-test.git/repo"
)

func GetLastProjects() (gitlabLastProjects DTO.GitlabLastProjects, err error) {
	gitlabNodes, errGet := repo.GetLastProjects()
	if errGet != nil {
		return gitlabLastProjects, errGet
	}
	for i, node := range gitlabNodes.Projects.Nodes {
		gitlabLastProjects.Names = gitlabLastProjects.Names + node.Name
		if i != len(gitlabNodes.Projects.Nodes)-1 {
			gitlabLastProjects.Names = gitlabLastProjects.Names + " - "
		}
		gitlabLastProjects.SumOfAllForks = gitlabLastProjects.SumOfAllForks + node.ForksCount
	}
	return gitlabLastProjects, nil
}
