package commands

import (
	"github.com/spf13/cobra"
)

// Root command for the application
var rootCmd = &cobra.Command{
	Use:   "testgen", 
	Short: "Testgen is a CLI tool for generating tests",
	Long:  "Testgen is a CLI tool that helps automate test generation for your codebase.",
}

func Execute() error {
	// Add flags and subcommands to the root cmd
	AddVersionFlag(rootCmd)
	AddHelpCommand(rootCmd) 

	// Root cmd default behavior 
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		// If no subcommands are provided, show the help message
	}

	// Execute the root command
	return rootCmd.Execute()
}
