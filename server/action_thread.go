package main

import (
	"time"

	"github.com/f7ed0/go-multiplayer-game/server/globals"
)

func actionThread() {
	t2 := time.Now()
	for true {
		globals.Players.Lock()
		for _, v := range globals.Players.Players {
			v.ApplyEvent(float32(time.Since(t2).Milliseconds()) / 1000)
		}
		t2 = time.Now()
		globals.Players.Unlock()
		time.Sleep(5 * time.Millisecond)
	}

}
