package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You must specify one argument as the file path.")
		os.Exit(1)
	}

	path := os.Args[1]
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

func charCount(line string) map[rune]int { // notice anything different here from previous exercises?
	m := make(map[rune]int, 0)
	for _, c := range line {
		m[c] = m[c] + 1
	}
	return m
}
