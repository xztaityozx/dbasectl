package encode

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

// EncodeはfileをBase64エンコードして返す
func Encode(file string) (string, error) {
	fp, err := os.Open(file)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(fp)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}
