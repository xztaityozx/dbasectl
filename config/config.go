package config

import (
	"path/filepath"
	"runtime"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type Config struct {
	// DocBaseのアクセストークン
	Token string
	// 所属しているチームの名前
	Name string
	// タイムアウト(msec)
	Timeout time.Duration
}

// Load はコンフィグファイルを読んで内容を Config 構造体にして返す
func Load(path string) (Config, error) {
	if path != "" {
		// パスが明示的に指定されていればそっちを使う
		viper.SetConfigFile(path)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			return Config{}, err
		}

		// linux/macOSは$HOME/.config/dbasectl/以下を見る
		viper.AddConfigPath(filepath.Join(home, ".config", "dbasectl"))
		// Windowsなら追加で $HOME\AppData\Roaming\dbasectlも見る
		if runtime.GOOS == "windows" {
			viper.AddConfigPath(filepath.Join(home, "AppData", "Roaming", "dbasectl"))
		}
		// ファイル名は dbasectl.{json,toml,yaml}など。viperが解釈できればなんでも
		viper.SetConfigName("dbasectl")
	}

	err := viper.ReadInConfig()

	return Config{Token: viper.GetString("Token"), Name: viper.GetString("Name"), Timeout: viper.GetDuration("Timeout")}, err
}
