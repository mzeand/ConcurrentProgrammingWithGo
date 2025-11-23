package main

//----------------------------------------
// grepdirrec.go
// A simple program that searches for a keyword in all files
// within a specified directory and its subdirectories using goroutines.
// Usage: go run grepdirrec.go keyword dirname
//----------------------------------------

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func grepfile(keyword, filename string) {
	defer wg.Done()
	fmt.Println("Reading file:", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	if strings.Contains(string(data), keyword) {
		fmt.Printf("File %s contains the keyword '%s'\n", filename, keyword)
	}
}

func searchWithWalkDir(keyword, rootDir string) error {
	return filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error accessing %s: %v\n", path, err)
			return nil
		}

		if !d.IsDir() {
			wg.Add(1)
			go grepfile(keyword, path)
		}
		return nil
	})
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: grepdirrec <keyword> <directory>")
		return
	}

	keyword := args[0]
	dir := args[1]

	fmt.Printf("Searching for keyword '%s' in directory '%s' recursively...\n", keyword, dir)

	err := searchWithWalkDir(keyword, dir)
	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		return
	}

	wg.Wait()
	fmt.Println("Search completed.")
}
