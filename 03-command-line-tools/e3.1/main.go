package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
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

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %s", err)
		os.Exit(1)
	}

	proverbs, err := loadProverbs(file)
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

func loadProverbs(file io.Reader) ([]*proverb, error) {
	proverbs := make([]*proverb, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		proverbs = append(proverbs, newProverb(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return proverbs, nil
}
