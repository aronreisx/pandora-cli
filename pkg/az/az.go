package az

import (
	"os/exec"

	"github.com/aronreisx/senses-cli/pkg/app"
)

type CreatePullRequestArgs struct {
	Title string
	ID string
}

func CreatePullRequest(title string, id string, sourceBranch string) string {
	resource := "senses"
	targetBranch := "master"

	cmd := exec.Command("az", "repos", "pr", "create", "--title", title, "--delete-source-branch", "--auto-complete", "--work-items", id, "-r", resource, "-s", sourceBranch, "-t", targetBranch)
	output, err := cmd.CombinedOutput()
	if err != nil {
		app.Error("")
	}
	return string(output)
}
