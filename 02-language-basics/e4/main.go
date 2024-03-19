package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("proverbs.txt")
	if err != nil {
		panic(fmt.Errorf("failed to open file: %s", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s\n", line)
		for k, v := range charCount(line) {
			fmt.Printf("'%s'=%d, ", k, v)
		}
		fmt.Print("\n\n")
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("failed to scan: %s", err))
	}
}

func charCount(line string) map[string]int {
	m := make(map[string]int, 0)
	for _, char := range line {
		m[string(char)] = m[string(char)] + 1
	}
	return m
}
