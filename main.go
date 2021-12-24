package main

import (
	"log"
	"os"
	"path/filepath"
)

func projectInit(projectType, projectPath string) {

	err := os.Mkdir(projectPath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	if projectType == "terraform" {
		terraformProject(projectPath)
	}
}

func terraformProject(projectPath string) {
	terraformFiles := [...]string{"main.tf", "variables.tf", "output.tf"}
	terraformDir := "terraform"

	err := os.Mkdir(filepath.Join(projectPath, terraformDir), 00755)
	if err != nil {
		log.Print(err)
	}

	for i := range terraformFiles {
		f, err := os.Create(filepath.Join(projectPath, terraformDir, terraformFiles[i]))
		if err != nil {
			log.Print(err)
		}

		f.Close()
	}
}

func main() {
	projectInit("terraform", "test_dir")

}
