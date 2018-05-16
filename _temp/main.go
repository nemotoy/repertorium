package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/sky0621/repertorium/pkg"
	"golang.org/x/net/context"
)

// TODO 機能実現スピード最優先での実装なので要リファクタ
func main() {
	// TODO env, argでまとめる
	var owner string
	eOwner := os.Getenv("OWNER")
	if eOwner == "" {
		fOwner := flag.String("o", "", "Target Owner")
		flag.Parse()
		owner = *fOwner
	} else {
		owner = eOwner
	}
	if owner == "" {
		fmt.Println("Please set 'Target Owner' by environment variable (key is OWNER) or flag (using '-o')")
		os.Exit(-1)
	}
	fmt.Printf("target owner is %s\n", owner)

	// TODO 全リポジトリ取得するまでページング（１ページ１００がMAXの様子）
	cli := pkg.NewGitHubClient()
	options := &github.RepositoryListOptions{ListOptions: github.ListOptions{Page: 1, PerPage: 500}}
	results, err := cli.GetRepositoriesList(context.Background(), owner, options)
	if err != nil {
		panic(err)
	}
	options2 := &github.RepositoryListOptions{ListOptions: github.ListOptions{Page: 2, PerPage: 500}}
	results2, err := cli.GetRepositoriesList(context.Background(), owner, options2)
	if err != nil {
		panic(err)
	}
	results = append(results, results2...)

	file, err := os.Create("./result.json")
	if err != nil {
		fmt.Println("02")
		panic(err)
	}
	defer file.Close()

	for _, result := range results {
		resultJSON, err := json.Marshal(&result)
		if err != nil {
			fmt.Println("03")
			panic(err)
		}
		file.Write(resultJSON)
		file.Write([]byte("\n"))
	}
}
