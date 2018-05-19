# repertorium
GitHubから指定オーナーのリポジトリを取得

★作成中★

## ■function

・フィルタリング（リポジトリ名の正規表現等）

・チェックアウトブランチ指定

## ■environment
go version
go version go1.9.4 linux/amd64

## ■ref

https://developer.github.com/v3/

https://developer.github.com/v3/repos/#list-user-repositories

## ■curl
<pre>
curl -i https://api.github.com/users/sky0621/repos?per_page=200 > /tmp/sky0621repos.txt
cat /tmp/sky0621repos.txt | grep clone_url
    "clone_url": "https://github.com/sky0621/aws-describe-prj.git",
    "clone_url": "https://github.com/sky0621/dockerfile-gowithdep.git",
    "clone_url": "https://github.com/sky0621/dockerfile-gowithglide.git",
    "clone_url": "https://github.com/sky0621/dockerfile-gowithgrpc.git",
    "clone_url": "https://github.com/sky0621/document.git",
    "clone_url": "https://github.com/sky0621/fons.git",
    "clone_url": "https://github.com/sky0621/fun-grpcgo.git",
    "clone_url": "https://github.com/sky0621/gitlab-seek-expert.git",
    "clone_url": "https://github.com/sky0621/go-algo.git",
　　　　・
　　　　・
　　　　・
    "clone_url": "https://github.com/sky0621/study-spring-rest.git",
    "clone_url": "https://github.com/sky0621/study-springboot.git",
    "clone_url": "https://github.com/sky0621/study-springboot-cli.git",
    "clone_url": "https://github.com/sky0621/study-springboot-maven.git",
    "clone_url": "https://github.com/sky0621/study-terraform.git",
</pre>

## ■MEMO

### add command
<pre>
cobra add parent
cobra add child -p parentCmd
</pre>
