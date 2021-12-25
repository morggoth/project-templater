package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing/object"
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

	addInitialCommit(projectPath)
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

func addInitialCommit(path string) {
	r, err := git.PlainOpen(path)
	errCheck(err)

	w, err := r.Worktree()
	errCheck(err)

	_, err = w.Add(".")
	errCheck(err)

	status, err := w.Status()
	errCheck(err)
	fmt.Println(status)

	// Get local git-config, merged with global config
	config, err := r.ConfigScoped(config.GlobalScope)
	errCheck(err)

	_, err = w.Commit("Initial commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  config.User.Name,
			Email: config.User.Email,
			When:  time.Now(),
		},
	})
	errCheck(err)
}

func main() {
	projectInit("terraform", "test_dir")
}
