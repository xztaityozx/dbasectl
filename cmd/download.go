package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/xztaityozx/dbasectl/request"
)

var downloadCmd = &cobra.Command{
	Use:     "download",
	Short:   "ファイルをダウンロードします",
	Long:    `IDを指定してDocBaseからファイルをダウンロードします。`,
	Example: "dbasectl download [Id]",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Idは1つ指定してください")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		name, _ := cmd.Flags().GetString("out")

		if len(name) == 0 {
			name = id
		}

		if err := downloadDo(id, name); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().StringP("out", "o", "", "出力ファイル名です。指定しない場合idが使われます。")
}

func downloadDo(id, name string) error {
	if len(id) == 0 {
		return fmt.Errorf("Idをに空できません")
	}

	req, err := request.New(cfg, http.MethodGet, request.Download)
	if err != nil {
		return err
	}

	if err := req.WithLogger(logger).SetLastPath(id).Build(); err != nil {
		return err
	}

	res, err := req.Do(ctx)
	if err != nil {
		if res == nil {
			return err
		}

		content, rErr := ioutil.ReadAll(res)
		if rErr != nil {
			return rErr
		}

		return fmt.Errorf("%v\n%s", err, string(content))
	}

	content, err := ioutil.ReadAll(res)

	return ioutil.WriteFile(name, content, 0644)
}
