package display

import (
	"encoding/gob"
	"net"
	"slices"
	"time"

	"github.com/f7ed0/go-multiplayer-game/client/handleplayer"
	"github.com/f7ed0/go-multiplayer-game/commons/entity/player"
	"github.com/f7ed0/golog/lg"
)

func (w *Window) Communication(conn net.Conn) {

	out := gob.NewEncoder(conn)
	in := gob.NewDecoder(conn)
	var msg string
	err := in.Decode(&msg)
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
	for !w.exit {
		err = in.Decode(&msg)
		if err != nil {
			lg.Error.Fatalln(err.Error())
			return
		}

		w.Me.RLock()
		err = out.Encode(&w.Me.ActionBuffer)
		err = out.Encode(&w.Me.Orientation)
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
		w.Me.Position = pcore.Position
		w.Me.LastTime = time.Now()
		w.Me.Unlock()

		var pcores []player.PlayerCore
		err = in.Decode(&pcores)
		if err != nil {
			lg.Error.Fatalln(err.Error())
			return
		}

		w.OtherMutex.Lock()
		showed := []string{}
		for _, item := range pcores {
			item.LastTime = time.Now()
			_, ok := w.Other[item.Hash]
			if ok {
				w.Other[item.Hash].PlayerCore = item
				w.Other[item.Hash].Here = true
			} else {
				p := handleplayer.FromPlayerCore(item)
				w.Other[item.Hash] = &p
			}
			showed = append(showed, item.Hash)

		}
		for k := range w.Other {
			if !slices.Contains(showed, k) {
				w.Other[k].Here = false
			}
		}
		w.OtherMutex.Unlock()

	}
}
