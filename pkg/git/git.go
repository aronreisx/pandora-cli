package git

import (
	"os/exec"
	"strings"
)

type GitCommandReturn struct {
	Output []byte
	Error  error
}

var baseCommand = "git"

type CreateBranchArgs struct {
	TaskType    string
	TeamName    string
	TaskId      string
	Description string
}

type CreateCommitMessageArgs struct {
	TaskType    string
	Description string
}

func CreateBranch(branchName string) GitCommandReturn {
	return commandPattern(baseCommand, "checkout", "-b", branchName)
}

func AddCommitAll(commitMessage string) GitCommandReturn {
	return commandPattern(baseCommand, "commit", "-am", commitMessage)
}

func PushUpstream(branchName string) GitCommandReturn {
	return commandPattern(baseCommand, "push", "-u", "origin", branchName)
}

func commandPattern(baseCommand string, args ...string) GitCommandReturn {
	cmd := exec.Command(baseCommand, args...)
	output, err := cmd.Output()
	return GitCommandReturn{output, err}
}

func CreateBranchName(createBranchArgs CreateBranchArgs) string {
	return (createBranchArgs.TaskType +
		"/" +
		createBranchArgs.TeamName +
		"-" +
		createBranchArgs.TaskId +
		createBranchArgs.getDescription())
}

func CreateCommitMessage(createCommitMessageArgs CreateCommitMessageArgs) string {
	commitType := commitTypeHandler(createCommitMessageArgs.TaskType)
	return (commitType + ": " + createCommitMessageArgs.Description)
}

func commitTypeHandler(taskType string) string {
	switch taskType {
	case "feature":
		return "feat"
	case "bugfix":
		return "fix"
	default:
		return taskType
	}
}

func (c CreateBranchArgs) getDescription() string {
	return strings.ReplaceAll(c.Description, " ", "-")
}
