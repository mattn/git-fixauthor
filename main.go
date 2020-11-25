package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const name = "git-fixauthor"

const version = "0.0.1"

var revision = "HEAD"

const fix = `
if [ "$GIT_COMMITTER_EMAIL" = "%[1]s" ]
then
    export GIT_COMMITTER_NAME="$[2]s"
    export GIT_COMMITTER_EMAIL="$[3]s"
fi
if [ "$GIT_AUTHOR_EMAIL" = "%s[1]s" ]
then
    export GIT_AUTHOR_NAME="$[2]s"
    export GIT_AUTHOR_EMAIL="$[3]s"
fi
`

var (
	fromEmail = flag.String("email", "", "old email")
)

func gitconfig(name string) string {
	b, err := exec.Command("git", "config", name).CombinedOutput()
	if err != nil {
		log.Fatalf("cannot get new name/email: %v", err)
	}
	return strings.TrimSpace(string(b))
}

func main() {
	flag.Parse()
	if *fromEmail == "" {
		flag.Usage()
		os.Exit(1)
	}
	name := gitconfig("user.name")
	email := gitconfig("user.email")
	cmd := exec.Command("git", "filter-branch", "-f", "--env-filter", fmt.Sprintf(fix, *fromEmail, name, email), "HEAD")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
