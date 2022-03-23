package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ip",
	Long:  `A cli that can track ip location`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	// Do Stuff Here
	//	action.TrackIP(args[0])
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
