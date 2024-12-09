package commands

import (
	"fmt"
	"github.com/Atrolide/go-testgen/internal/misc"
	"github.com/spf13/cobra"
	"os"
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

	// Add the --version flag to the root command
	rootCmd.Flags().BoolP("version", "v", false, "Print the version of testgen")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
