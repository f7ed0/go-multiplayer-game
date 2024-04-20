package globals

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"github.com/f7ed0/go-multiplayer-game/commons/entity/player"
)

var Players = NewPlayerArray()

type PlayerArray struct {
	Players map[string]*player.PlayerCore
	sync.RWMutex
}

func (pa *PlayerArray) AddNewPlayer() string {
	pa.Lock()
	defer pa.Unlock()
	hasher := md5.New()
	a := fmt.Sprintf("%v", time.Now().Unix())
	h := hasher.Sum([]byte(a))
	hash := hex.EncodeToString(h)
	pa.Players[hash] = new(player.PlayerCore)
	*pa.Players[hash] = player.NewPlayer()
	return hash
}

func (pa *PlayerArray) DropPlayer(hash string) {
	pa.Lock()
	defer pa.Unlock()
	delete(pa.Players, hash)
}

func NewPlayerArray() PlayerArray {
	return PlayerArray{
		Players: make(map[string]*player.PlayerCore),
	}
}

func (pa *PlayerArray) GiveOmitMe(meID string) (ret []player.PlayerCore) {
	pa.Lock()
	defer pa.Unlock()
	ret = []player.PlayerCore{}
	for k, v := range pa.Players {
		if k == meID {
			continue
		}
		ret = append(ret, *v)
	}
	return
}
