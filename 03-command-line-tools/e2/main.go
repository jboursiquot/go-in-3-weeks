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

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Failed to read file: %s", err)
		os.Exit(1)
	}
	proverbs := string(bs)

	lines := strings.Split(proverbs, "\n")
	for _, l := range lines {
		fmt.Printf("%s\n", l)
		for k, v := range charCount(l) {
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

func charCount(line string) map[rune]int {
	m := make(map[rune]int, 0)
	for _, c := range line {
		m[c] = m[c] + 1
	}
	return m
}
