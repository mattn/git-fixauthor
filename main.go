package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var (
	fromName = flag.String("name", "", "old name")
)

func gitconfig(name string) string {
	b, err := exec.Command("git", "config", name).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(b))
}

func main() {
	flag.Parse()
	if *fromName == "" {
		flag.Usage()
		os.Exit(1)
	}
	name := gitconfig("user.name")
	email := gitconfig("user.email")
	cmd := exec.Command("git", "filter-branch", "-f", "--env-filter",
		fmt.Sprintf(strings.Join([]string{
			`if [ "$GIT_AUTHOR_NAME" = "%[1]s" ];`,
			`then`,
			`GIT_AUTHOR_NAME="%[2]s";`,
			`GIT_AUTHOR_EMAIL="%[3]s";`,
			`fi`}, "\n"),
			*fromName, name, email), "HEAD")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
