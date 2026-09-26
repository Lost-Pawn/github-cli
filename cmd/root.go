package cmd

import (
	"fmt"
	"github-cli/internals/activity"
)

func RootCmd(args []string) error {
	switch args[0] {
	case "github-activity":
		err := getUserName(args[1:])
		if err != nil {
			return err
		}
	case "git-watch":
		// err := activity.WatchRepository(args[1:])
		// if err != nil {
		// 	return err
		// }
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