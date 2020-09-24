package request

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/xztaityozx/dbasectl/encode"
)

const (
	Upload EndPoint = "attachments"
)

type content struct {
	name    string
	content string
}

type body []content

// FileUpload は files を DocBaseにアップロードする
func (r Request) FileUpload(ctx context.Context, files ...string) error {
	dict := map[string]string{}
	b := body{}

	for _, v := range files {
		fi, err := os.Stat(v)
		if err != nil {
			return err
		}

		// ディレクトリはUploadできないので弾く
		if fi.IsDir() {
			return errors.New(fmt.Sprintf("%s is directory", v))
		}

		// /dev/null みたいなスペシャルファイルはUploadできない
		if !fi.Mode().IsRegular() {
			return errors.New(fmt.Sprintf("%s is not regular file", v))
		}

		data, err := encode.Encode(v)
		if err != nil {
			return err
		}

		b = append(b, content{name: fi.Name(), content: data})
		dict[fi.Name()] = v
	}

	//json, err := json.Marshal(b)
	//if err != nil {
	//return err
	//}

	//url, err := r.getUrl(Upload)
	//if err != nil {
	//return err
	//}

	//req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(json))
	//if err != nil {
	//return err
	//}

	//req.Header.Set("Content-Type", "application/json")
	//r.setAccessToken(req)

	//client := &http.Client{}
	//res, err := client.Do(req)
	//if err != nil {
	//return err
	//}

	//defer res.Body.Close()

	return nil
}
