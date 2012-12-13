package goit_test

import "fmt"
import "github.com/traviscline/goit"

func ExampleInitRepository() {
	// create a new non-bare repository
	repo, err := goit.InitRepository("/tmp/testing.git", false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(repo.Workdir())
}
