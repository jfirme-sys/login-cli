package main

import (
	"envsecret-cli/commands"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "envsecret",
	Short: "A CLI tool to inject environment variables",
}

func main() {
	rootCmd.AddCommand(commands.LoginCmd)
	rootCmd.AddCommand(commands.FetchProjectsCmd)
	rootCmd.AddCommand(commands.SelectProjectCmd)
	rootCmd.AddCommand(commands.InjectCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
