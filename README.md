# dbasectl
DocBase API cli client

DocBaseのコンソールクライアントです。チーム名とそのチーム内で有効なAPITokenを使って、メモの検索、更新、ファイルアップロードやダウンロードができます。

# Install
## go getする
```zsh
$ go get -u github.com/xztaityozx/dbasectl
```

## バイナリをGitHub Releasesからダウンロードする
TODO

# Configuration
`dbasectl`はチーム名やAPITokenの設定にJSONやYAMLといったファイルを使用します。規定のパスはWindows/Linux/macOSすべてで、`$HOME/.config/dbasectl/dbasectl.{json,yaml,toml}`です。

Windowsの場合のみ、`$HOME/AppData/Roaming/dbasectl/dbasectl.{json,yaml,toml}`が利用できます。

## 設定項目
| 項目 | 説明 | 例 | デフォルト値 |
|:--:|:--:|:--:|:--:|
|`token`(required)|DocBaseAPIを利用するために必要なAPITokenです。 詳しくは[DocBaseのAPIリファレンス](https://help.docbase.io/posts/45703)を読んで下さい| - | - |
|`name`(required)|所属しているチームの名前です| team | - |
|`timeout`|APIアクセスの制限時間(nsec)。負数を指定すると、無限になります| 100 | -1 |

## Example
### JSON
```json
{
  "token": "API-TOKEN",
  "name": "sugoi-team"
}
```

### YAML
```yaml
token: API-TOKEN
name: sugoi-team
timeout: 100000
```

# SubCommands
dbasectlはいくつかのサブコマンドで機能を切り替えます

|機能名|サブコマンド名|状態|  
|:--:|:--:|:--:|  
|ファイルアップロード|upload|ok|  
|ファイルダウンロード|download|ok|  
|メモ検索|search|wip|  
|メモ更新|update|todo|  
|メモ投稿|post|todo|  
|メモ削除|delete|todo|  
|メモのアーカイブ・アーカイブ解除|archive|todo|  
|コメント投稿・削除|comment|todo|  
|タグの取得|tag|todo|  
|グループの作成・検索・追加・削除|group|todo|  