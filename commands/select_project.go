package commands

import (
	"envsecret-cli/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var SelectProjectCmd = &cobra.Command{
	Use:   "select-project",
	Short: "Select a project",
	Run: func(cmd *cobra.Command, args []string) {
		projectId, _ := cmd.Flags().GetString("projectId")
		err := utils.SaveSelectedProject(projectId)
		if err != nil {
			fmt.Println("Error selecting project:", err)
			return
		}

		fmt.Println("Project selected:", projectId)
	},
}

func init() {
	SelectProjectCmd.Flags().String("projectId", "", "Project ID")
	SelectProjectCmd.MarkFlagRequired("projectId")
}
