package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("proverbs.txt")
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}
	proverbs := string(bs)

	lines := strings.Split(proverbs, "\n")
	for _, l := range lines {
		fmt.Printf("%s\n", l)
		for k, v := range charCount(l) {
			fmt.Printf("'%s'=%d, ", k, v)
		}
		fmt.Print("\n\n")
	}
}

func charCount(line string) map[string]int {
	m := make(map[string]int, 0)
	for _, char := range line {
		m[string(char)] = m[string(char)] + 1
	}
	return m
}
