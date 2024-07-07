package commands

import (
	"envsecret-cli/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var FetchProjectsCmd = &cobra.Command{
	Use:   "fetch-projects",
	Short: "Fetch available projects",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		projects, err := utils.FetchProjects(token)
		if err != nil {
			fmt.Println("Error fetching projects:", err)
			return
		}

		fmt.Println("Projects:", projects)
	},
}

func init() {
	FetchProjectsCmd.Flags().String("token", "", "API token")
	FetchProjectsCmd.MarkFlagRequired("token")
}
