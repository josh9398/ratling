package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version_short = fmt.Sprintf("print the version number of %s", binary)
	version_long  = fmt.Sprintf("all software has versions. This is %s's", binary)
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: version_short,
	Long:  version_long,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("version %s\ngit commit %s\nbuild date %s\n", version, commit, date)
		return nil
	},
}
