package commands

import (
	"fmt"
	"github.com/mitchellh/colorstring"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "testgen",
	Short: "Testgen is a sample test file generator",
	Long:  `A fast generator for sample test files, supporting multiple programming languages.`,
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			// Call versionCmd to print the version
			versionCmd.Run(cmd, args)
			return
		}
		// No flag outputs the help message
		// TODO: Enhance error handling
		if err := cmd.Help(); err != nil {
			errorMessage := fmt.Sprintf("\n[bold][red]Error: %s\n", err)
			fmt.Println(colorstring.Color(errorMessage))
			return
		}
	},
}

func init() {
	// Disable the default completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Disable command suggestions for unknown commands
	rootCmd.DisableSuggestions = true

	// Silence default usage and error outputs
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true

	// Add the --version flag to the root command
	rootCmd.Flags().BoolP("version", "v", false, "Print the version of testgen")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// TODO: Enhance error handling
		errorMessage := fmt.Sprintf("\n[bold][red]Error: %s\n", err)
		fmt.Println(colorstring.Color(errorMessage))
		return
	}
}
