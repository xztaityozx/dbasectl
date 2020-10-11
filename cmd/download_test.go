package cmd

import (
	"github.com/stretchr/testify/assert"
	"github.com/xztaityozx/dbasectl/config"
	"testing"
)

func Test_downloadDo(t *testing.T) {
	as := assert.New(t)
	t.Run("Idを空にしたとき", func(t *testing.T) {
		as.Error(downloadDo("", ""), "エラーが返されるべき")
	})

	t.Run("ファイル名を空にしたとき", func(t *testing.T) {
		as.Error(downloadDo("id", ""), "エラーが返されるべき")
	})

	t.Run("コンフィグが正しくないとき", func(t *testing.T) {
		cfg = config.Config{}
		as.Error(downloadDo("id", "name"), "エラーが返されるべき")
	})
}
