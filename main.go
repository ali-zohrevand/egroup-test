package main

import (
	"fmt"

	"github.com/ali-zohrevand/egroup-test.git/service"
)

func main() {
	name, numberOfForks, err := service.GetLastProjects()
	if err != nil {
		panic(err)
	}
	fmt.Println("Last Projects in Gitlab")
	fmt.Println("names: ", name)
	fmt.Println("sum of all forks: ", numberOfForks)
}
