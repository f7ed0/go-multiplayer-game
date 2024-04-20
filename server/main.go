package main

import (
	"log"
	"net"

	"github.com/f7ed0/go-multiplayer-game/commons/lg"
)

const (
	HOSTIP   = "0.0.0.0"
	HOSTPORT = "44444"
)

func main() {
	lg.Init(lg.ALL, true)
	lg.Info.Println("Starting server...")
	server, err := net.Listen("tcp", HOSTIP+":"+HOSTPORT)
	if err != nil {
		log.Default().Fatal("Erreur a l'initialisation :" + err.Error())
	}
	defer server.Close()
	go actionThread()
	for {
		newplayer, err := server.Accept()
		if err != nil {
			log.Default().Fatal("Error Accepting new client : " + err.Error())
		}
		lg.Info.Println("Someone trying to connect...")
		playerThread(newplayer)
	}

}
