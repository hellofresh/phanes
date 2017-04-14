package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	providerName string
	providerURL  string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "phanes",
	Short: "This is the client generator for an identity provider",
	Long: `This is a simple, CLI tool that helps you to create oauth clients 
	for many different identity providers. Through this tool you can create and delete clients, 
	you can also hook this up in any web application that you wish. At the moment we support the hellofresh, 
	google and facebook identity providers.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.phanes.yaml)")
	RootCmd.PersistentFlags().StringVarP(&providerName, "provider", "p", "hellofresh", "The authentication provider to create the oauth client")
	RootCmd.PersistentFlags().StringVarP(&providerURL, "provider-url", "u", "hellofresh", "The identity provider client endpoint")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".phanes") // name of config file (without extension)
	viper.AddConfigPath("$HOME")   // adding home directory as first search path
	viper.AutomaticEnv()           // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
