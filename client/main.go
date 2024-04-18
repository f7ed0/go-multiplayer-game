package main

import (
	"github.com/f7ed0/go-multiplayer-game/client/display"
	"github.com/f7ed0/go-multiplayer-game/commons/lg"
)

func main() {
	lg.Init(lg.ALL, true)
	w, err := display.NewWindow()
	if err != nil {
		lg.Error.Fatalln("Error initialising window : ", err.Error())
	}
	err = w.LoadMap("assets/maps/test1")
	if err != nil {
		lg.Error.Fatalln("Error loading map : ", err.Error())
	}
	w.MainLoop()
}
