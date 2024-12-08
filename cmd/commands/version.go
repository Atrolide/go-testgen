package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Testgen",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Pass the version dynamically from env
		// TODO: Output platform data, e.g., linux_amd64
		fmt.Println("Testgen v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
