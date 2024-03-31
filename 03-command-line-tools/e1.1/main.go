package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("You must specify one argument as the file path.")
		os.Exit(1)
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file: %s", err)
		os.Exit(1)
	}
	defer file.Close()

	mime, err := mimeType(file)
	if err != nil {
		fmt.Printf("Failed to determine MIME type: %s", err)
		os.Exit(1)
	}
	if mime != "text/plain; charset=utf-8" {
		fmt.Printf("Unsupported file type: %s.\n", mime)
		os.Exit(1)
	}

	bs, err := io.ReadAll(file)
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

// mimeType returns the media/MIME type of the file based on the first 512 bytes of the file.
func mimeType(file io.ReadSeeker) (string, error) {
	// Why 512 bytes? Because the magic number for most file types is within the first 512 bytes.
	// What is a magic number? It's a unique sequence of bytes that identifies a file type.
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// Reset the read pointer to the beginning of the file
	// otherwise subsequent read operations will start from the 513th byte.
	// Try commenting this out and see what happens!
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", fmt.Errorf("failed to seek file: %w", err)
	}

	return http.DetectContentType(buffer), nil
}

func charCount(line string) map[rune]int { // notice anything different here from previous exercises?
	m := make(map[rune]int, 0)
	for _, c := range line {
		m[c] = m[c] + 1
	}
	return m
}
