package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// VERSION is set during build
	VERSION string
	// ConfigFile can be set for command
	ConfigFile string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "instago",
	Short: "Print list of most liked photos from instagram page",
}

// Execute adds all child commands to the root command
func Execute(version string) {
	VERSION = version

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&ConfigFile, "config", "", "config file (default is $HOME/.instago.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if ConfigFile != "" {
		viper.SetConfigFile(ConfigFile)
	}

	viper.SetConfigName(".cli")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
