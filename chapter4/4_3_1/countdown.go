package main

import (
	"fmt"
	"sync"
	"time"
)

func countDown(seconds *int, mutex *sync.Mutex) {
	for *seconds > 0 {
		time.Sleep(1 * time.Second)
		mutex.Lock()
		*seconds -= 1
		mutex.Unlock()
	}
}

func main() {
	mutex := sync.Mutex{}
	count := 5
	go countDown(&count, &mutex)

	for count > 0 {
		time.Sleep(500 * time.Millisecond)
		mutex.Lock()
		fmt.Println(count)
		mutex.Unlock()
	}
}
