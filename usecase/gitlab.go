package usecase

import (
	"github.com/ali-zohrevand/egroup-test.git/model/DTO"
	"github.com/ali-zohrevand/egroup-test.git/repo"
)

//GetLastProjects After receiving the names of the projects and their specifications in this layer, we want to perform
//the requested action on that data. Any other similar processing and operation can be done in this layer.
func GetLastProjects() (gitlabLastProjects DTO.GitlabLastProjects, err error) {
	gitlabNodes, errGet := repo.GetLastProjects()
	if errGet != nil {
		return gitlabLastProjects, errGet
	}
	for i, node := range gitlabNodes.Projects.Nodes {
		gitlabLastProjects.Names = gitlabLastProjects.Names + node.Name
		//This check is necessary so that we do not add a dash after the last name.
		if i != len(gitlabNodes.Projects.Nodes)-1 {
			gitlabLastProjects.Names = gitlabLastProjects.Names + " - "
		}
		gitlabLastProjects.SumOfAllForks = gitlabLastProjects.SumOfAllForks + node.ForksCount
	}
	return gitlabLastProjects, nil
}
