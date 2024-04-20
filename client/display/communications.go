package display

import (
	"encoding/gob"
	"log"
	"net"

	"github.com/f7ed0/go-multiplayer-game/commons/entity/player"
	"github.com/f7ed0/go-multiplayer-game/commons/lg"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
)

func (w *Window) Communication() {
	conn, err := net.Dial("tcp", "localhost:44444")
	if err != nil {
		log.Fatal("Error during connection : " + err.Error())
	}
	defer conn.Close()
	out := gob.NewEncoder(conn)
	in := gob.NewDecoder(conn)
	var msg string
	err = in.Decode(&msg)
	if err != nil {
		lg.Error.Fatalln(err.Error())
		return
	}
	lg.Debug.Println(msg)
	err = out.Encode(&msg)
	if err != nil {
		lg.Error.Fatalln(err.Error())
		return
	}
	lg.Info.Println("Connected to", conn.RemoteAddr())
	for {
		err = in.Decode(&msg)
		if err != nil {
			lg.Error.Fatalln(err.Error())
			return
		}

		w.Me.RLock()
		err = out.Encode(&w.Me.ActionBuffer)
		w.Me.RUnlock()

		if err != nil {
			lg.Error.Fatalln(err.Error())
			return
		}
		var pcore player.PlayerCore
		err := in.Decode(&pcore)
		if err != nil {
			lg.Error.Fatalln(err.Error())
			return
		}

		w.Me.Lock()
		if objects.Diff(w.Me.Position, pcore.Position).N2_2D() > 30 {
			lg.Debug.Println("ROLBACKED")
			w.Me.Position = pcore.Position
		}
		w.Me.Unlock()

		var pcores []player.PlayerCore
		err = in.Decode(&pcores)
		if err != nil {
			lg.Error.Fatalln(err.Error())
			return
		}

		w.OtherMutex.Lock()
		w.Other = pcores
		w.OtherMutex.Unlock()

	}
}
