package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/xztaityozx/dbasectl/output"
	"github.com/xztaityozx/dbasectl/request"
	"github.com/xztaityozx/dbasectl/result"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "メモ、ユーザー、グループの検索を行います",
	Long: `メモ、ユーザー、グループを検索します。デフォルトはメモ検索です。
検索には 'author:name' や 'group:name' といったDocBaseでも使える検索が可能です。

--include-user-groupsは--userでユーザー検索をするときのみ有効です。またチームのオーナーや管理者のみ利用可能です。
dbasectlでは、この機能をデフォルトでOFFにしています。ONにするには、設定ファイルに 'ActivateOwnerFeature' を追加してください

--user-searchと--group-searchは同時に指定できません。
`,
	Run: func(cmd *cobra.Command, args []string) {
		us, _ := cmd.Flags().GetBool("user-search")
		iug, _ := cmd.Flags().GetBool("include-user-groups")
		page, _ := cmd.Flags().GetInt("page")
		pp, _ := cmd.Flags().GetInt("per-page")
		of, _ := cmd.Flags().GetString("output")
		gs, _ := cmd.Flags().GetBool("group-search")

		s := search{
			user:         us,
			includeGroup: iug,
			page:         page,
			perPage:      pp,
			output:       output.New(of),
			query:        args,
			group:        gs,
		}

		if err := s.do(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	searchCmd.Flags().Bool("user-search", false, "ユーザー検索を行います")
	searchCmd.Flags().Bool("include-user-groups", false, "ユーザー検索時、所属するグループのリストも一緒に返すようにします。オーナー、管理者のみ利用可能です")
	searchCmd.Flags().IntP("page", "P", 1, "取り出すデータ群のページ番号です")
	searchCmd.Flags().Int("per-page", 100, "一度に取り出すデータの個数です。API制限により、最大値は200です。")
	searchCmd.Flags().String("output", "json", "出力形式です(json, yaml, text)")
	searchCmd.Flags().Bool("group-search", false, "グループ検索を行います")

	rootCmd.AddCommand(searchCmd)
}

type search struct {
	user         bool
	group        bool
	includeGroup bool
	page         int
	perPage      int
	output       output.Format
	query        []string
}

func (s search) verify() error {
	if s.includeGroup {
		if !s.user {
			return fmt.Errorf("--include-user-groupオプションは、ユーザー検索モード(--user)時のみ有効です")
		}

		if !cfg.ActivateOwnerFeature {
			return fmt.Errorf("オーナー機能が有効になっていません。")
		}
	}

	if s.user && s.group {
		return fmt.Errorf("ユーザー検索とグループ検索は同時にできません")
	}

	if s.page == 0 {
		return fmt.Errorf("0番目のページ番号は存在しません")
	}

	if s.perPage > 200 || s.perPage < 1 {
		return fmt.Errorf("--per-pageの値域は 1 <= x <= 200 です")
	}

	if s.output != output.Yaml && s.output != output.Json && s.output != output.Text {
		return fmt.Errorf("%s は非対応のフォーマットです", s.output)
	}

	return nil
}

func (s search) do() error {
	if err := s.verify(); err != nil {
		return err
	}

	req, err := func() (request.Request, error) {
		if s.user {
			return request.New(cfg, http.MethodGet, request.UserSearch)
		} else if s.group {
			return request.New(cfg, http.MethodGet, request.GroupSearch)
		} else {
			return request.New(cfg, http.MethodGet, request.PostSearch)
		}
	}()

	if err != nil {
		return err
	}

	req.WithLogger(logger).
		AddParameter("page", fmt.Sprint(s.page)).
		AddParameter("per_page", fmt.Sprint(s.perPage))

	if s.includeGroup {
		req.AddParameter("include_user_groups", "true")
	}

	if len(s.query) != 0 {
		key := "q"
		if s.group {
			key = "name"
		}

		req.AddParameter(key, strings.Join(s.query, " "))
	}

	if err := req.Build(); err != nil {
		return err
	}

	res, err := req.Do(ctx)
	if err != nil {
		return err
	}
	defer res.Close()

	data, err := ioutil.ReadAll(res)
	if err != nil {
		return err
	}

	r, err := func() (result.Stringer, error) {
		if s.group {
			var rt result.Groups
			err := json.Unmarshal(data, &rt)
			return rt, err
		} else if s.user {
			var rt result.Users
			err := json.Unmarshal(data, &rt)
			return rt, err
		}
		var rt result.Posts
		err := json.Unmarshal(data, &rt)
		return rt, err
	}()

	return s.output.Print(r, os.Stdout)
}
