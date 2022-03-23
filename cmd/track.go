package cmd

import (
	"cli-ip-tracker/action"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(trackCmd)
}

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Show the details of IP",
	Long:  `Show details of IP`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			action.TrackIP(args[0])
		}
	},
}