package commands

import (
	"envsecret-cli/utils"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var InjectCmd = &cobra.Command{
	Use:   "inject",
	Short: "Inject environment variables into the current shell",
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")

		secrets, err := utils.FetchSecrets(token)
		if err != nil {
			fmt.Println("Error fetching secrets:", err)
			return
		}

		for key, value := range secrets {
			os.Setenv(key, value)
		}

		if len(args) > 0 {
			command := exec.Command(args[0], args[1:]...)
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
			command.Stdin = os.Stdin
			err = command.Run()
			if err != nil {
				fmt.Println("Error executing command:", err)
			}
		} else {
			fmt.Println("No command provided to run with the injected environment variables")
		}
	},
}

func init() {
	InjectCmd.Flags().String("token", "", "API token for fetching secrets")
	InjectCmd.MarkFlagRequired("token")
}
