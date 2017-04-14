package cmd

import (
	"github.com/fatih/color"
	"github.com/hellofresh/phanes/pkg/provider"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes an oauth client by ID",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		provider, err := provider.Create(providerName, providerURL)
		if err != nil {
			color.Red(err.Error())
			return
		}

		id, err := cmd.Flags().GetString("name")
		if err != nil || id == "" {
			color.Red("Please provide the id of the client")
			return
		}

		err = provider.Remove(id)
		if err != nil {
			color.Red(err.Error())
			return
		}

		color.Green("Client %s was removed!", id)
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
	removeCmd.Flags().String("id", "", "The client's ID")
}
