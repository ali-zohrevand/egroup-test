package service

import "github.com/ali-zohrevand/egroup-test.git/usecase"

//GetLastProjects In this service, we only want to get the name and number of forks to display after execution.
//We can have other services to provide json output in Rest APIs.
func GetLastProjects() (names string, numberOfForks int, err error) {
	lastProjects, errGet := usecase.GetLastProjects()
	if errGet != nil {
		return names, numberOfForks, errGet
	}
	return lastProjects.Names, lastProjects.SumOfAllForks, nil
}
