package cmd

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func Test_Do(t *testing.T) {
	as := assert.New(t)

	baseDir := filepath.Join(os.TempDir(), "dbasectl", "test")
	as.Nil(os.MkdirAll(baseDir, 0755))

	t.Run("ディレクトリはUploadできない", func(t *testing.T) {
		as.Error(do(baseDir))
	})

	t.Run("スペシャルファイルはUploadできない", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip()
		}

		as.Error(do("/dev/null"), "/dev/nullはUpできない")
		as.Error(do("/dev/nu??"), "Globも展開できる")
	})

	t.Run("存在しないファイルはUploadできない", func(t *testing.T) {
		p := filepath.Join(baseDir, "none")
		as.Error(do(p))
	})

	t.Run("指定が0個のときはUploadできない", func(t *testing.T) {
		as.Error(do())
	})

	t.Run("base64エンコーディング出来ないファイルを指定した", func(t *testing.T) {
	})

	_ = os.RemoveAll(baseDir)
}
