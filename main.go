package main

import (
	"fmt"

	"github.com/ali-zohrevand/egroup-test.git/service"
	"github.com/ali-zohrevand/egroup-test.git/util"
)

func main() {
	standardLogger := util.NewLogger()
	name, numberOfForks, err := service.GetLastProjects()
	if err != nil {
		standardLogger.Panic(err)
		panic(err)
	}
	fmt.Println("=======\nLast Projects in Gitlab")
	fmt.Println("names: ", name)
	fmt.Println("sum of all forks: ", numberOfForks)
}
