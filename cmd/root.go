package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/xztaityozx/dbasectl/config"
)

var rootCmd = &cobra.Command{
	Use:   "dbasectl",
	Short: "CLI tool for DocBase API",
	Long:  `CLI tool for DocBase API.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello")
	},
}

var cfgFile string
var cfg config.Config

func init() {
	var err error
	cfg, err = config.Load(cfgFile)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load config")
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
	}
}
