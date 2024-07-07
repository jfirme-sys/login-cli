package commands

import (
	"envsecret-cli/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var SelectProjectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select a project",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		err := utils.SaveSelectedProject(id)
		if err != nil {
			fmt.Println("Error selecting project:", err)
			return
		}

		fmt.Println("Project selected:", id)
	},
}

func init() {
	SelectProjectCmd.Flags().String("id", "", "Project ID")
	SelectProjectCmd.MarkFlagRequired("id")
}
