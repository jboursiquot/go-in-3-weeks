package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	path := pathFromFlag()
	if path == "" {
		path = pathFromEnv()
	}

	if path == "" {
		fmt.Println("You must specify one the file path with -f or as FILE environment variable.")
		os.Exit(1)
	}

	proverbs, err := loadProverbs(path)
	if err != nil {
		fmt.Printf("Failed to load proverbs: %s", err)
		os.Exit(1)
	}

	for _, p := range proverbs {
		fmt.Printf("%s\n", p.line)
		for k, v := range p.chars {
			fmt.Printf("'%c'=%d, ", k, v)
		}
		fmt.Print("\n\n")
	}
}

func pathFromFlag() string {
	path := flag.String("f", "", "file flag")
	flag.Parse()
	return *path
}

func pathFromEnv() string {
	return os.Getenv("FILE")
}

func loadProverbs(path string) ([]*proverb, error) {
	var proverbs []*proverb

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bs), "\n")
	for _, line := range lines {
		proverbs = append(proverbs, newProverb(line))
	}

	return proverbs, nil
}
