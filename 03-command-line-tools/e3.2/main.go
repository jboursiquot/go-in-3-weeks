package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const proverbs = `Don't communicate by sharing memory, share memory by communicating.
Concurrency is not parallelism.
Channels orchestrate; mutexes serialize.
The bigger the interface, the weaker the abstraction.
Make the zero value useful.
interface{} says nothing.
Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.
A little copying is better than a little dependency.
Syscall must always be guarded with build tags.
Cgo must always be guarded with build tags.
Cgo is not Go.
With the unsafe package there are no guarantees.
Clear is better than clever.
Reflection is never clear.
Errors are values.
Don't just check errors, handle them gracefully.
Design the architecture, name the components, document the details.
Documentation is for users.
Don't panic.`

func main() {
	fr := strings.NewReader(proverbs)

	proverbs, err := loadProverbs(fr)
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
