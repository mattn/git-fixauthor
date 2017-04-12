package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func gitconfig(name string) string {
	b, err := exec.Command("git", "config", name).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(b))
}

func main() {
	name := gitconfig("user.name")
	email := gitconfig("user.email")
	cmd := exec.Command("git", "filter-branch", "-f", "--env-filter",
		fmt.Sprintf(""+
			"GIT_AUTHOR_NAME='%[1]s';"+
			"GIT_AUTHOR_EMAIL='%[2]s';"+
			"GIT_COMMITTER_NAME='%[1]s';"+
			"GIT_COMMITTER_EMAIL='%[2]s';"+
			"HEAD",
			name, email))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
