package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload file to docbase",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("This is upload sub command")
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
