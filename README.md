# repertorium
GitHubから指定オーナーのリポジトリを取得（ `git clone` ）します。

リポジトリ取得後、設定ファイル（ `$HOME/.repertorium.yaml` ）記載の指定ブランチをチェックアウトします。

取得先に既に同一リポジトリが存在する場合は、 `git pull` を行います。

## ■require
- git

## ■function

- フィルタリング
  - リポジトリ名の正規表現指定
  - 対象言語指定

- チェックアウトブランチ指定

## ■prepare
- download binary

https://github.com/sky0621/repertorium/releases

- download config

https://github.com/sky0621/repertorium/blob/master/.repertorium.yaml

- edit config

- set config to [$HOME/.repertorium.yaml]

## ■exec(binary) f.e. Linux
$ ./repertorium_linux_amd64 get
