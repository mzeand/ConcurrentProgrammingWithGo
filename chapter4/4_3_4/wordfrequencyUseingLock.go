package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func countWords(url string, frequency map[string]int, m *sync.Mutex) {
	fmt.Println("Fetching URL:", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()
	var word string
	for {
		_, err := fmt.Fscan(resp.Body, &word)
		if err != nil {
			break
		}
        m.Lock()
		frequency[word]++
        m.Unlock()
	}

}

func main() {
	var frequency = make(map[string]int)
	m := sync.Mutex{}

	for i := 1000; i < 1010; i++ {
		url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
		go countWords(url, frequency, &m)
	}

	time.Sleep(1 * time.Second)
	for word, freq := range frequency {
		if freq > 0 {
			fmt.Printf("%s -> %d\n", word, freq)
		}
	}

}
