package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/xztaityozx/dbasectl/encode"
	"github.com/xztaityozx/dbasectl/request"
	"net/http"
	"os"
	"path/filepath"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload file to docbase",
	Run: func(cmd *cobra.Command, args []string) {
		if err := do(args...); err != nil {
			logrus.Fatal(err)
		}
	},
}

type content struct {
	name    string
	content string
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}

func do(files ...string) error {
	dic := map[string]string{}

	for _, file := range files {
		// Globかもしれんので展開する
		paths, err := filepath.Glob(file)
		if err != nil {
			return err
		}

		// 展開したあとのファイルリストをそれぞれStatする
		for _, p := range paths {
			fp, err := os.Stat(p)
			if err != nil {
				return err
			}

			// ディレクトリはUpできない
			if fp.IsDir() {
				return fmt.Errorf("ディレクトリ(%s)はアップロードできません", p)
			}

			// スペシャルファイルはUpできない
			if !fp.Mode().IsRegular() {
				return fmt.Errorf("スペシャルファイル(%s)はアップロードできません", p)
			}

			dic[fp.Name()] = p
		}
	}

	// Upload候補が0個だった
	if len(dic) == 0 {
		return fmt.Errorf("アップロードすべきファイルがありません")
	}

	var b []content
	for name, p := range dic {
		encoded, err := encode.Encode(p)
		if err != nil {
			// base64エンコード出来なかった
			return err
		}

		b = append(b, content{name: name, content: encoded})
	}

	// json文字列へ
	jsonBytes, err := json.Marshal(b)
	if err != nil {
		return err
	}

	req, err := request.New(cfg, http.MethodPost, request.Upload)
	if err != nil {
		return err
	}

	err = req.WithLogger(logger).SetBody(bytes.NewBuffer(jsonBytes)).Build()
	if err != nil {
		return err
	}

	// リクエストする
	res, err := req.Do(ctx)
	if err != nil {
		return err
	}
	defer res.Close()
	return PrintJson(res)
}
