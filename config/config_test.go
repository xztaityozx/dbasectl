package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func Test_Load(t *testing.T) {
	as := assert.New(t)
	cfg := Config{
		Token:                "token",
		Name:                 "name",
		Timeout:              100,
		ActivateOwnerFeature: false,
	}

	home, _ := homedir.Dir()

	t.Run("load on linux/macOS/windows", func(t *testing.T) {
		baseDir := filepath.Join(home, ".config", "dbasectl")
		_ = os.MkdirAll(baseDir, 0755)
		for _, v := range []string{"json", "yaml", "toml"} {
			path := filepath.Join(baseDir, fmt.Sprintf("dbasectl.%s", v))
			var data []byte
			if v == "json" {
				data, _ = json.Marshal(cfg)
			} else if v == "yaml" {
				data, _ = yaml.Marshal(cfg)
			} else if v == "toml" {
				data, _ = toml.Marshal(cfg)
			}

			as.Nil(ioutil.WriteFile(path, data, 0644))

			{
				res, err := Load(path)
				as.Nil(err)
				as.Equal(cfg, res, "パスを指定してロードできる")
			}

			{
				res, err := Load("")
				as.Nil(err)
				as.Equal(cfg, res, "パスを指定しなくてもロードできる")
			}
		}

		_ = os.RemoveAll(baseDir)
	})

	t.Run("load on windows", func(t *testing.T) {
		if runtime.GOOS != "windows" {
			t.Skip()
		}

		baseDir := filepath.Join(home, "AppData", "Roaming", "dbasectl")
		_ = os.MkdirAll(baseDir, 0755)
		for _, v := range []string{"json", "yaml", "toml"} {
			path := filepath.Join(baseDir, fmt.Sprintf("dbasectl.%s", v))
			var data []byte
			if v == "json" {
				data, _ = json.Marshal(cfg)
			} else if v == "yaml" {
				data, _ = yaml.Marshal(cfg)
			} else if v == "toml" {
				data, _ = toml.Marshal(cfg)
			}

			as.Nil(ioutil.WriteFile(path, data, 0644))

			{
				res, err := Load(path)
				as.Nil(err)
				as.Equal(cfg, res, "パスを指定してロードできる")
			}

			{
				res, err := Load("")
				as.Nil(err)
				as.Equal(cfg, res, "パスを指定しなくてもロードできる")
			}
		}

		_ = os.RemoveAll(baseDir)
	})
}
