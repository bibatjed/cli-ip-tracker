package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cli-ip-tracker",
	Short: "Hugo is a very fast static site generator",
	Long:  `A cli that can track ip location`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	// Do Stuff Here
	//	fmt.Println("bithcass")
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
