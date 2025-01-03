package commands

import (
	"fmt"
	"github.com/Atrolide/go-testgen/internal/generator"
	"github.com/mitchellh/colorstring"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:     "generate",
	Aliases: []string{"g", "gen"},
	Short:   "Generate sample test file(s)",
	Args:    cobra.NoArgs,
	Example: "testgen generate --file <file_path>",
	Run: func(cmd *cobra.Command, args []string) {
		fileFlag, _ := cmd.Flags().GetString("file")
		if fileFlag != "" {
			// Generate a test file for a given file
			generator.FileGenByFile(fileFlag)
			return
		}
		// TODO: Enhance error handling
		if err := cmd.Help(); err != nil {
			errorMessage := fmt.Sprintf("\n[bold][red]Error: %s\n", err)
			fmt.Println(colorstring.Color(errorMessage))
			return
		}
	},
}

func init() {
	// Add the --file flag to the generate command
	generateCmd.Flags().StringP("file", "f", "", "Generate a test for a given file")

	rootCmd.AddCommand(generateCmd)
}
