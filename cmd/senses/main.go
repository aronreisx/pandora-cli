package main

import (
	"os"
	"strings"

	"github.com/aronreisx/go-intro/pkg/app"
	"github.com/aronreisx/go-intro/pkg/az"
	"github.com/aronreisx/go-intro/pkg/git"
)

type Option string

type flags struct {
	description string
	taskType    string
	team        string
	id          string
}

const (
	Workflow Option = "workflow"
)

func main() {
	app.MinimumArgsAmountValidator(1)
	option := Option(os.Args[1])
	optionHandler(option)
}

func optionHandler(option Option) {
	parsedFlags := parseFlags(os.Args[2:])
	flags := flags{
		description: parsedFlags["description"],
		taskType:    parsedFlags["type"],
		team:        parsedFlags["team"],
		id:          parsedFlags["id"],
	}

	createBranchArgs := git.CreateBranchArgs{
		TaskType:    flags.taskType,
		TeamName:    flags.team,
		TaskId:      flags.id,
		Description: flags.description,
	}

	createCommitMessageArgs := git.CreateCommitMessageArgs{
		TaskType:    flags.taskType,
		Description: flags.description,
	}

	switch option {
	case Workflow:
		branchName := git.CreateBranchName(createBranchArgs)
		commitMessage := git.CreateCommitMessage(createCommitMessageArgs)
		git.CreateBranch(branchName)
		git.AddCommitAll(commitMessage)
		git.PushUpstream(branchName)
		az.CreatePullRequest(commitMessage, flags.id, branchName)
		app.Success()
	default:
		app.Error(app.HelpMessage)
	}
}

func parseFlags(args []string) map[string]string {
	flags := make(map[string]string)
	for i, arg := range args {
		if strings.HasPrefix(arg, "--") {
			flag := strings.ToLower(strings.TrimPrefix(arg, "--"))
			lastIndex := len(args) - 1
			nextIndex := i + 1
			if i < lastIndex && !strings.HasPrefix(args[nextIndex], "--") {
				flags[flag] = args[nextIndex]
			} else {
				app.Error(app.HelpMessage)
			}
		}
	}
	return flags
}
