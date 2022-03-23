package cmd


import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of IP tracker",
	Long:  `Show version of IP Tracker`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ip-tracker v1.0")
	},
}