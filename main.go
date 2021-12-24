package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

func errCheck(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func projectInit(projectType, projectPath string) {

	err := os.Mkdir(projectPath, 0755)
	errCheck(err)

	initGitRepo(projectPath)

	if projectType == "terraform" {
		terraformProject(projectPath)
	}
}

func terraformProject(projectPath string) {
	terraformFiles := [...]string{"main.tf", "variables.tf", "output.tf"}
	terraformDir := "terraform"

	err := os.Mkdir(filepath.Join(projectPath, terraformDir), 00755)
	errCheck(err)

	for i := range terraformFiles {
		f, err := os.Create(filepath.Join(projectPath, terraformDir, terraformFiles[i]))
		errCheck(err)

		f.Close()
	}
}

func initGitRepo(path string) {
	_, err := git.PlainInit(path, false)
	errCheck(err)
}

func main() {
	projectInit("terraform", "test_dir")

}
