package main

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
	dir := args[1]
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			go grepfile(keyword, dir+"/"+entry.Name())
		}
	}
	runtime.Gosched()
	time.Sleep(1 * time.Second)
	fmt.Println("Main function completed.")
}
