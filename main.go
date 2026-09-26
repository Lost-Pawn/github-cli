package main

import (
	"fmt"
	"os"
	"github-cli/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`Please provide a command.`)
		return

	}

	cmd := cmd.RootCmd(os.Args[1:]) 
	if err := cmd; err != nil {
		fmt.Println("Error:", err)
	}
}