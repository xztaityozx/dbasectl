package output

import (
	"encoding/json"
	"fmt"
	"github.com/xztaityozx/dbasectl/result"
	"gopkg.in/yaml.v2"
	"io"
)

// 出力形式
type Format int

const (
	Text Format = iota
	Json
	Yaml
)

// String はFormat型を文字列にする
func (f Format) String() string {
	return []string{"text", "json", "yaml"}[f]
}

func (f Format) Print(r result.Stringer, w io.Writer) error {
	data, err := func() ([]byte, error) {
		if f == Json {
			return json.MarshalIndent(r, "", "  ")
		} else if f == Yaml {
			return yaml.Marshal(r)
		} else if f == Text {
			return []byte(r.String()), nil
		}

		return nil, fmt.Errorf("%v は未対応のフォーマットです", f)
	}()
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}
