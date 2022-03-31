package service

import "github.com/ali-zohrevand/egroup-test.git/usecase"

func GetLastProjects() (names string, numberOfForks int, err error) {
	lastProjects, errGet := usecase.GetLastProjects()
	if errGet != nil {
		return names, numberOfForks, errGet
	}
	return lastProjects.Names, lastProjects.SumOfAllForks, nil
}
