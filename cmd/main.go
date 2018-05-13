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

	cli := pkg.NewGitHubClient()
	// TODO set perpage
	options := &github.RepositoryListOptions{}
	results, err := cli.GetRepositoriesList(context.Background(), owner, options)
	if err != nil {
		panic(err)
	}

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
