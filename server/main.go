package main

import "github.com/f7ed0/go-multiplayer-game/commons/lg"

func main() {
	lg.Init(lg.ALL, true)
	lg.Info.Println("Starting server...")
}
