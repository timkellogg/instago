package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show heft.io client version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(RootCmd.Use + " v" + VERSION)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
