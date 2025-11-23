package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func matchReader(matchEvents *[]string, mutex *sync.RWMutex) {
	for i := 0; ; i++ {
		if mutex.TryLock() {
			*matchEvents = append(*matchEvents, "Match event"+strconv.Itoa(i))
			mutex.Unlock()
			fmt.Println("Append match event", i)
		} else {
			fmt.Println("Failed to acquire lock, skipping event", i)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func clientHandler(mEvents *[]string, mutex *sync.RWMutex, st time.Time) {
	for i := 0; i < 100; i++ {
		mutex.RLock()
		allEvents := copyAllEvents(mEvents)
		mutex.RUnlock()
		timeTaken := time.Since(st)
		time.Sleep(200 * time.Millisecond)
		fmt.Println(len(allEvents), "events copied in", timeTaken)
	}
}

func copyAllEvents(matchEvents *[]string) []string {
	allEvents := make([]string, 0, len(*matchEvents))
	for _, event := range *matchEvents {
		allEvents = append(allEvents, event)
	}
	return allEvents
}

func main() {
	mutex := sync.RWMutex{}
	var matchEvents = make([]string, 0, 10000)
	for j := 0; j < 10000; j++ {
		matchEvents = append(matchEvents, "Match event")
	}

	go matchReader(&matchEvents, &mutex)
	startTime := time.Now()
	go clientHandler(&matchEvents, &mutex, startTime)
	time.Sleep(100 * time.Second)
}
