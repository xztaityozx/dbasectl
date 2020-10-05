package encode

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Encode(t *testing.T) {
	dir := filepath.Join(os.TempDir(), "dbasectl_test")
	as := assert.New(t)
	as.Nil(os.MkdirAll(dir, 0755))
	file := filepath.Join(dir, "test")

	defer os.RemoveAll(dir)

	t.Run("ファイルがオープンできなかったとき", func(t *testing.T) {
		_, err := Encode(file)
		as.NotNil(err, "エラーが投げられるべき")
	})

	t.Run("ファイルが有るとき", func(t *testing.T) {
		data := []byte("This is test data")
		as.Nil(ioutil.WriteFile(file, data, 0644))

		expect := base64.StdEncoding.EncodeToString(data)
		actual, err := Encode(file)

		as.Nil(err)
		as.Equal(expect, actual, "正常にBase64エンコードできる")
	})
}
