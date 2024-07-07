package main

import (
	"envsecret-cli/commands"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected '--login', '--get-projects' or '--inject-secrets' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "--login":
		commands.Login(os.Args[2], os.Args[3])
	case "--get-projects":
		commands.GetProjects(os.Args[2])
	case "--inject-secrets":
		commands.InjectSecrets()
	default:
		fmt.Println("Expected '--login', '--get-projects' or '--inject-secrets' subcommands")
		os.Exit(1)
	}
}
