package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// flags
	force bool
)

func init() {
	rootCmd.AddCommand(pruneCmd)
	sendCmd.Flags().BoolVarP(&force, "force", "y", false, "force prune cache")
}

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: `prune local cache`,
	Long:  `large files are chunked & stored in a cache, this command prunes that cache`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("prunning...")
	},
}
