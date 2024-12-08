package commands

import (
	"github.com/Atrolide/go-testgen/internal/generator"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a file",
	Run: func(cmd *cobra.Command, args []string) {
		generator.FileGen()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
