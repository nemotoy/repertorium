# repertorium
GitHubから指定オーナーのリポジトリを取得（ `git clone` ）します。

リポジトリ取得後、設定ファイル（ `$HOME/.repertorium.yaml` ）記載の指定ブランチをチェックアウトします。

取得先に既に同一リポジトリが存在する場合は、 `git pull` を行います。

## ■require
- go
- dep
- git

## ■function

- フィルタリング
  - リポジトリ名の正規表現指定
  - 対象言語指定

- チェックアウトブランチ指定

## ■environment
$ go version

go version go1.9.4 linux/amd64

## ■exec
$ dep ensure

$ go run main.go --config .repertorium.yaml get
