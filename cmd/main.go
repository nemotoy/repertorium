package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/net/context"
)

func main() {
	cli := github.NewClient(nil)
	results, res, err := cli.Repositories.List(context.Background(), "sky0621", nil)
	if err != nil {
		fmt.Println("01")
		panic(err)
	}
	if res != nil {
		fmt.Printf("%#v\n", res)
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
	}
}
