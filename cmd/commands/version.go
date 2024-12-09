package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Testgen",
	Run: func(cmd *cobra.Command, args []string) {

		/*
			TEMP solution!
			TODO: Define a custom version command:
				- Pass the version dynamically from env
			 	- Output platform data, e.g., linux_amd64
		*/

		fmt.Println("Testgen v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
