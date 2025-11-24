package main

import (
	"fmt"
	"sync"
	"time"
)

func playerHandler(cond *sync.Cond, playerRemaining *int, playerID int, gameStarted *bool) {
	cond.L.Lock()
	fmt.Println(playerID, ": Connected")
	*playerRemaining--
	if *playerRemaining == 0 && !*gameStarted {
		fmt.Println("All players connected! Starting game immediately.")
		*gameStarted = true
		cond.Broadcast()
	}
	for *playerRemaining > 0 && !*gameStarted {
		fmt.Println(playerID, ": Waiting for more players")
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Println(playerID, ": Game started! Ready player", playerID)
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	playerRemaining := 4
	gameStarted := false

	go func() {
		time.Sleep(3 * time.Second)
		cond.L.Lock()
		if !gameStarted {
			fmt.Println("Timer: 3 seconds elapsed, starting game with available players!")
			gameStarted = true
			playerRemaining = 0
			cond.Broadcast()
		}
		cond.L.Unlock()
	}()

	for playerID := 1; playerID <= 4; playerID++ {
		go playerHandler(cond, &playerRemaining, playerID, &gameStarted)
		time.Sleep(1 * time.Second)
	}

	time.Sleep(10 * time.Second)
}
