package main

import (
	"encoding/gob"
	"net"
	"time"

	"github.com/f7ed0/go-multiplayer-game/commons/entity/player"
	"github.com/f7ed0/go-multiplayer-game/commons/lg"
	"github.com/f7ed0/go-multiplayer-game/server/globals"
)

func playerThread(conn net.Conn) {
	out := gob.NewEncoder(conn)
	in := gob.NewDecoder(conn)
	msg := "OK"
	err := out.Encode(&msg)
	if err != nil {
		lg.Error.Println(err.Error())
		return
	}
	err = in.Decode(&msg)
	if err != nil {
		lg.Error.Println(err.Error())
		return
	}
	if msg == "OK" {
		lg.Info.Println("Auth Successfull")
		var acts player.ActionBuffer
		id := globals.Players.AddNewPlayer()
		defer globals.Players.DropPlayer(id)

		globals.Players.RLock()
		lg.Info.Println(len(globals.Players.Players))
		globals.Players.RUnlock()
		for {
			err = out.Encode("GO")
			if err != nil {
				lg.Error.Println(err.Error())
				return
			}
			err = in.Decode(&acts)
			globals.Players.Lock()
			value := globals.Players.Players[id]
			value.ActionBuffer = acts
			globals.Players.Unlock()
			if err != nil {
				lg.Error.Println(err.Error())
				return
			}

			globals.Players.RLock()
			out.Encode(value)
			globals.Players.RUnlock()

			out.Encode(globals.Players.GiveOmitMe(id))

			time.Sleep(50 * time.Millisecond)
		}
	}
}
