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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello")
	},
}

var cfgFile string
var cfg config.Config

func init() {
	var err error

	rootCmd.Flags().StringVar(&cfgFile, "config", "", "path to config file")
	rootCmd.Flags().String("token", "", "Your access token for docbase api")
	rootCmd.Flags().String("name", "n", "Name of your docbase team")

	cfg, err = config.Load(cfgFile)
	if err != nil {
		logrus.WithError(err).Fatal("Failed to load config")
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
	}
}
