// cmd/commands/help.go
package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var customHelpMessage = `Testgen is a CLI tool that helps automate test generation for your codebase.

Usage:
  testgen [flags] [args]
  testgen [command] [args]

Available commands:
  version    Show the current version of Testgen
  help       Displays this output

Available flags:
  -v, --version    Show the current version of Testgen
  -h, --help       Displays this output
`

// Adds the --help flag to the root command and the help subcommand
func AddHelpCommand(rootCmd *cobra.Command) {
	// Override Cobra's default Help function to show custom help
	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		// Print the custom help message when --help or -h is used
		fmt.Println(customHelpMessage)
	})

	// Add the help subcommand 
	helpCmd := &cobra.Command{
		Use:   "help",
		Short: "Show help message",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(customHelpMessage)
			os.Exit(0)
		},
	}

	// Add the help subcommand to the root command
	rootCmd.AddCommand(helpCmd)
}
