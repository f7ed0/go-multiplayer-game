package main

import (
	"time"

	"github.com/f7ed0/go-multiplayer-game/server/globals"
)

func actionThread() {
	for {
		globals.Players.Lock()
		for _, v := range globals.Players.Players {
			v.ApplyEvent(float32(time.Since(v.LastTime).Milliseconds()) / 1000)
			v.LastTime = time.Now()
		}
		globals.Players.Unlock()
		time.Sleep(5 * time.Millisecond)
	}

}
