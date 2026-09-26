package main

import (
	"fmt"
	"os"
	"github-cli/cmd"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(`Please provide a command and username as an argument.`)
		return

	}

	cmd := cmd.RootCmd(os.Args[1:]) // use [2:0]
	if err := cmd; err != nil {
		fmt.Println("Error:", err)
	}
}