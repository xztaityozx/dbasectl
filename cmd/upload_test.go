package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xztaityozx/dbasectl/config"
)

func Test_Do(t *testing.T) {
	as := assert.New(t)

	baseDir := filepath.Join(os.TempDir(), "dbasectl", "test")
	as.Nil(os.MkdirAll(baseDir, 0755))

	t.Run("ディレクトリはUploadできない", func(t *testing.T) {
		as.Error(uploadDo(baseDir))
	})

	t.Run("スペシャルファイルはUploadできない", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip()
		}

		as.Error(uploadDo("/dev/null"), "/dev/nullはUpできない")
		as.Error(uploadDo("/dev/nu??"), "Globも展開できる")
	})

	t.Run("存在しないファイルはUploadできない", func(t *testing.T) {
		p := filepath.Join(baseDir, "none")
		as.Error(uploadDo(p))
	})

	t.Run("指定が0個のときはUploadできない", func(t *testing.T) {
		as.Error(uploadDo())
	})

	t.Run("configが不足しててUploadできない", func(t *testing.T) {
		p := filepath.Join(baseDir, "file")
		as.Nil(ioutil.WriteFile(p, []byte("それ"), 0644))

		cfg = config.Config{}
		as.Error(uploadDo(p))
		as.Error(uploadDo(filepath.Join(baseDir, "fi??")), "Globでも良い")
		_ = os.Remove(p)
	})

	_ = os.RemoveAll(baseDir)
}
