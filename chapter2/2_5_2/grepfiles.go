package main

//----------------------------------------
// grepfiles.go
// A simple program that searches for a keyword in files
// provided as command-line arguments using goroutines.
// Usage: go run grepfiles.go keyword file1.txt file2.txt
//----------------------------------------
import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func grepfile(keyword, filename string) {
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

func main() {
	args := os.Args[1:]
	keyword := args[0]
	fmt.Println("Searching for keyword:", keyword)
	files := args[1:]
	for _, filename := range files {
		go grepfile(keyword, filename)
	}
	runtime.Gosched()
	time.Sleep(1 * time.Second)
	fmt.Println("Main function completed.")
}
