package commands

import (
	"envsecret-cli/utils"
	"fmt"

	"github.com/spf13/cobra"
)

var FetchSecretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "Fetch secrets for the selected project",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := utils.LoadToken()
		if err != nil {
			fmt.Println("Error loading token:", err)
			return
		}

		projectId, err := utils.LoadSelectedProject()
		if err != nil {
			fmt.Println("Error loading selected project:", err)
			return
		}

		secrets, err := utils.FetchSecrets(token, projectId)
		if err != nil {
			fmt.Println("Error fetching secrets:", err)
			return
		}

		for key, value := range secrets {
			fmt.Printf("%s: %s\n", key, value)
		}
	},
}
