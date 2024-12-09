package commands

import (
	"fmt"
	"os"

	"github.com/Atrolide/go-testgen/internal/misc"
	"github.com/mitchellh/colorstring"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "testgen",
	Short: "Testgen is a sample test file generator",
	Long:  `A fast generator for sample test files, supporting multiple programming languages.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the --version flag is set
		if versionFlag, _ := cmd.Flags().GetBool("version"); versionFlag {
			// Call versionCmd to print the version
			versionCmd.Run(cmd, args)
			os.Exit(0) // Exit after displaying the version
		}

		// If --version flag is not set, output ASCII art and help
		misc.AsciiArt() // Outputs the ASCII art
		cmd.Help()      // Outputs the help message
	},
}

func init() {
	// Disable the default completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// Hack: Disable command suggestions for unknown commands
	rootCmd.DisableSuggestions = true

	// Silence usage and errors
	rootCmd.SilenceUsage = true
	rootCmd.SilenceErrors = true

	// Add the --version flag to the root command
	rootCmd.Flags().BoolP("version", "v", false, "Print the version of testgen")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// Output error
		// Todo: Customize error message, kick out to another file, call a command here
		fmt.Println(colorstring.Color(fmt.Sprintf("\n[bold][red]Error: %s\n", err.Error())))

		// Output help
		// Todo: Output help for command on which the error occured
		rootCmd.Help()
		os.Exit(1)
	}
}
