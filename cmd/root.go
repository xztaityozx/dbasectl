package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
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

func init() {
	cobra.OnInitialize(initConfig)

	// サブコマンドまで使えるオプションたち
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "コンフィグファイルへのパス")
	rootCmd.PersistentFlags().String("token", "", "DocBase APIにアクセスするためのAPI Token")
	rootCmd.PersistentFlags().String("name", "", "DocBaseのチーム名")
	rootCmd.PersistentFlags().DurationP("timeout", "t", -1, "リクエストをタイムアウトする秒数(nsec)。負数で無限")
	rootCmd.PersistentFlags().Bool("verbose", false, "ログを出力しながら実行します")
	rootCmd.PersistentFlags().BoolP("pretty-print", "p", false, "レスポンスとして帰ってきたJSONを成形して出力します")

	// コンフィグとオプションのバインド
	for _, name := range []string{"token", "name", "timeout"} {
		if err := viper.BindPFlag(name, rootCmd.PersistentFlags().Lookup(name)); err != nil {
			logrus.WithError(err).Warn("failed to bind ", name, " option")
		}
	}
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
	loggerFormatter := &logrus.TextFormatter{ForceColors: true, TimestampFormat: time.RFC3339}

	if viper.GetBool("verbose") {
		logger = logrus.New()
		logger.SetOutput(os.Stderr)
		logger.SetFormatter(loggerFormatter)
	}

	logrus.SetFormatter(loggerFormatter)

	// 全体で使うコンテキスト
	var cancelFunc context.CancelFunc
	ctx, cancelFunc = context.WithCancel(context.Background())

	// Ctrl-Cでキャンセルできる
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)
	go func() {
		if logger != nil {
			logger.Warn("Ctrl+Cでキャンセルできます")
		}
		<-sigCh
		cancelFunc()
	}()

	if err := rootCmd.Execute(); err != nil {
	}
}

// PrintJson はviperのフラグを見て成形してSTDOUTに出したりする
func PrintJson(r io.Reader) error {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	if b, _ := rootCmd.PersistentFlags().GetBool("pretty-print"); b {
		var indented []byte
		r := bytes.NewBuffer(indented)
		_ = json.Indent(r, content, "", "  ")
		fmt.Println(r.String())
	} else {
		fmt.Println(string(content))
	}

	return nil
}
