package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version_short = fmt.Sprintf("Print the version number of %s", appName)
	version_long  = fmt.Sprintf("All software has versions. This is %s's", appName)
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
