package cmd

import (
	"github.com/fatih/color"
	"github.com/hellofresh/phanes/pkg/provider"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new oauth client",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		provider, err := provider.Create(providerName, providerURL)
		if err != nil {
			color.Red(err.Error())
			return
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil || name == "" {
			color.Red("Please provide the name of the client")
			return
		}

		redirectURL, _ := cmd.Flags().GetString("redirect-url")
		client, err := provider.Create(name, redirectURL)
		if err != nil {
			color.Red(err.Error())
			return
		}

		color.Green("Credentials for %s created!", name)
		color.Cyan("Client ID: %s", client.GetID())
		color.Cyan("Client Secret: %s", client.GetSecret())
	},
}

func init() {
	RootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "The client name")
	createCmd.Flags().StringP("redirect-url", "r", "http://localhost", "The callback url")
}
