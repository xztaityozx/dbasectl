package cmd

import (
	"context"
	"os"
	"time"

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
var logger *logrus.Logger
var ctx context.Context
var cancelFunc context.CancelFunc

func init() {
	cobra.OnInitialize(initConfig)

	// サブコマンドまで使えるオプションたち
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "コンフィグファイルへのパス")
	rootCmd.PersistentFlags().String("token", "", "DocBase APIにアクセスするためのAPI Token")
	rootCmd.PersistentFlags().String("name", "", "DocBaseのチーム名")
	rootCmd.PersistentFlags().DurationP("timeout", "t", -1, "リクエストをタイムアウトする秒数(msec)。負数で無限")
	rootCmd.PersistentFlags().Bool("verbose", false, "ログを出力しながら実行します")

	// コンフィグとオプションのバインド
	for _, name := range []string{"token", "name", "timeout"} {
		if err := viper.BindPFlag(name, rootCmd.PersistentFlags().Lookup(name)); err != nil {
			logrus.WithError(err).Warn("failed to bind ", name, " option")
		}
	}

	// 全体で使うコンテキスト
	ctx, cancelFunc = context.WithCancel(context.Background())
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
	if viper.GetBool("verbose") {
		logger = logrus.New()
		logger.SetOutput(os.Stderr)
		logger.Formatter = &logrus.TextFormatter{ForceColors: true, TimestampFormat: time.RFC3339}
	}

	if err := rootCmd.Execute(); err != nil {
	}
}
