package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "v0.0.1-alpha0"

var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("App Version:" + version)
	},
}

func init() {
	rootCmd.AddCommand(versionCommand)
}
