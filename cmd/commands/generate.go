package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a file",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Call generator method once created
		fmt.Println("Generating a file...")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
