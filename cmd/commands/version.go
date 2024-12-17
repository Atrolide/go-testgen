package commands

import (
	"github.com/Atrolide/go-testgen/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of Testgen",
	Run: func(cmd *cobra.Command, args []string) {

		version.ReadVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
