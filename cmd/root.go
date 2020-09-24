package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xztaityozx/dbasectl/config"
)

var rootCmd = &cobra.Command{
	Use:   "dbasectl",
	Short: "CLI tool for DocBase API",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info(cfg)
	},
}

var cfgFile string
var cfg config.Config

func init() {
	cobra.OnInitialize(initConfig)

	// サブコマンドまで使えるオプションたち
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "path to config file")
	rootCmd.PersistentFlags().String("token", "", "Your access token for docbase api")
	rootCmd.PersistentFlags().String("name", "", "Name of your docbase team")
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("name", rootCmd.PersistentFlags().Lookup("name"))
}

// initConfig は コンフィグのロードなどを行う
func initConfig() {
	var err error
	cfg, err = config.Load(cfgFile)

	if err != nil {
		logrus.WithError(err).Fatal("Failed to load config")
	}
}

// Execute はこのアプリのエントリーポイント
func Execute() {
	if err := rootCmd.Execute(); err != nil {
	}
}
