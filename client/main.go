package main

import (
	"log"
	"net"

	"github.com/f7ed0/go-multiplayer-game/client/display"
	"github.com/f7ed0/go-multiplayer-game/commons/lg"
)

func main() {
	lg.Init(lg.ALL, true)

	conn, err := net.Dial("tcp", "192.168.179.84:44444")
	if err != nil {
		log.Fatal("Error during connection : " + err.Error())
	}
	defer conn.Close()
	w, err := display.NewWindow()
	if err != nil {
		lg.Error.Fatalln("Error initialising window : ", err.Error())
	}
	err = w.LoadMap("assets/maps/test1")
	if err != nil {
		lg.Error.Fatalln("Error loading map : ", err.Error())
	}
	go w.Communication(conn)
	w.MainLoop()
}
