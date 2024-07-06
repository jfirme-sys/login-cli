package main

import (
	"envsecret-cli/commands"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected 'configure', 'get-secret' or 'inject-secrets' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "configure":
		commands.Configure()
	case "get-secret":
		commands.GetSecret(os.Args[2])
	case "inject-secrets":
		commands.InjectSecrets()
	default:
		fmt.Println("Expected 'configure', 'get-secret' or 'inject-secrets' subcommands")
		os.Exit(1)
	}
}
