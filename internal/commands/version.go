package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"runtime" 
)

// Todo: Pass version dynamically, based on GH release
var Version = "1.0.0" 

// Adds the --version flag to the root command and the version subcommand
func AddVersionFlag(rootCmd *cobra.Command) {
	// --version flag
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print the version of the app")

	// version subcommand
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version of the app",
		Run: func(cmd *cobra.Command, args []string) {
			// Get the platform and architecture information from the runtime package
			platform := fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH)

			// Print version and platform details
			fmt.Printf("Testgen Version: %s\n", Version)
			fmt.Printf("on %s\n", platform)

			os.Exit(0)
		},
	}

	// Add the version subcommand to the root command
	rootCmd.AddCommand(versionCmd)

	// Use PersistentPreRun to check if the version flag is set
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			// Get the platform and architecture information
			platform := fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH)

			// Print version along with platform details
			fmt.Printf("Testgen Version: %s\n", Version)
			fmt.Printf("Platform: %s\n", platform)

			// Exit after printing the version
			os.Exit(0)
		}
	}
}
