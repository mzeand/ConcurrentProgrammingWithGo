package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func catfile(filename string) {
	fmt.Println("Reading file:", filename)
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Printf("Contents of %s:\n%s\n", filename, data)
}

func main() {
	args := os.Args[1:]
	for _, filename := range args {
		go catfile(filename)
	}
	runtime.Gosched()
	time.Sleep(1 * time.Second)
	fmt.Println("Main function completed.")
}
