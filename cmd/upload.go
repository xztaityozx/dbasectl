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

		if err := do(args...); err != nil {
			logrus.Fatal(err)
		}
	},
}

type content struct {
	name    string
	content string
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}

func do(files ...string) error {
	return nil
}
