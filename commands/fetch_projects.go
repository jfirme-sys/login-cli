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
		token, err := utils.LoadToken()
		if err != nil {
			fmt.Println("Error loading token:", err)
			return
		}

		projects, err := utils.FetchProjects(token)
		if err != nil {
			fmt.Println("Error fetching projects:", err)
			return
		}

		for _, project := range projects {
			fmt.Printf("ID: %s, Name: %s\n", project.ID, project.Name)
		}
	},
}
