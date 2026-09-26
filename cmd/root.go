package cmd

import (
	"fmt"
	"os"
)

func RootCmd(args []string) error {
	switch args[1] {
	case "github-activity":
		
	case "git-watch":
		fmt.Println("watch repo")
	default:
		fmt.Println(`Invalid command. Please use 'github-activity' or 'git-watch'.`)
		fmt.Println("Example:")
		fmt.Println("> github-activity lost-pawn")
	}
	return nil
}

