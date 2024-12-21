package commands

import (
	"github.com/Atrolide/go-testgen/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v", "ver"},
	Short:   "Print the version of Testgen",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		version.ReadVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
