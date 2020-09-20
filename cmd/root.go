package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docbase",
	Short: "CLI tool for DocBase API",
	Long:  `CLI tool for DocBase API.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
	}
}
