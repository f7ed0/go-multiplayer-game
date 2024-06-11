package handleplayer

import "github.com/f7ed0/go-multiplayer-game/commons/entity/player"

type DispPlayer struct {
	player.PlayerCore
	FrameCount int
	Here       bool
}

func FromPlayerCore(p player.PlayerCore) DispPlayer {
	return DispPlayer{
		PlayerCore: p,
		FrameCount: 0,
		Here:       true,
	}
}
