package cmd

import (
	"flag"
	"fmt"
	"github-cli/internals"
)

func RootCmd(args []string) error {
	switch args[0] {
	case "github-activity":
		err := getUserName(args[1:])
		if err != nil {
			return err
		}
	case "git-watch":
		watchCmd := flag.NewFlagSet("git-watch", flag.ExitOnError)
		repo := watchCmd.String("repo", "", "GitHub repository in the format 'owner/repo'")
		interval := watchCmd.Int("interval", 60, "Interval in seconds to check for new commits")

		watchCmd.Parse(args[1:])

		if *repo == "" {
			return fmt.Errorf("Please provide a GitHub repository using the -repo flag.")
		}

		if *interval <= 0 {
			return fmt.Errorf("Please provide a valid interval in seconds using the -interval flag.")
		}

		fmt.Printf("Watching repository '%s' for new commits every %d seconds...\n", *repo, *interval)
		err := activity.WatchGitHubRepo(*repo, *interval)
		if err != nil {
			return err
		}
	default:
		fmt.Println(`Invalid command. Please use 'github-activity' or 'git-watch'.`)
		fmt.Println("Example:")
		fmt.Println("> github-activity lost-pawn")
	}
	return nil
}

func getUserName(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("Please provide a GitHub username as an argument.")
	}

	username := args[0]
	fmt.Println("Fetching GitHub activities for user:", username)

	activities, err := activity.FetchGitHubActivities(username)
	if err != nil {
		return err
	}

	return activity.DisplayActivities(username, activities)
}