package drawableentity

import (
	"github.com/f7ed0/go-multiplayer-game/client/handleplayer"
	"github.com/f7ed0/go-multiplayer-game/client/handleplayer/camera"
	"github.com/f7ed0/go-multiplayer-game/commons/entity"
	"github.com/f7ed0/go-multiplayer-game/commons/entity/player"
	"github.com/f7ed0/go-multiplayer-game/commons/lg"
	"github.com/veandco/go-sdl2/sdl"
)

var FrameCounter int = 0

func Draw(ent *handleplayer.DispPlayer, r *sdl.Renderer, cam camera.Camera, pm map[entity.EntityType]EntityModel, delta int) error {
	boundaries := r.GetViewport()
	offseted := cam.ApplyOffset(ent.Position, int(boundaries.W), int(boundaries.H))

	r.SetDrawColor(255, 100, 100, 100)
	var rec sdl.Rect
	walking := true
	if ent.ActionBuffer.Get(player.MOVE_UP) && !ent.ActionBuffer.Get(player.MOVE_DOWN) {
		rec = pm[entity.KNIGHT].States["walking"].GetFrame("up", ent.FrameCount)
	} else if ent.ActionBuffer.Get(player.MOVE_DOWN) && !ent.ActionBuffer.Get(player.MOVE_UP) {
		rec = pm[entity.KNIGHT].States["walking"].GetFrame("down", ent.FrameCount)
	} else if ent.ActionBuffer.Get(player.MOVE_LEFT) && !ent.ActionBuffer.Get(player.MOVE_RIGHT) {
		rec = pm[entity.KNIGHT].States["walking"].GetFrame("left", ent.FrameCount)
	} else if ent.ActionBuffer.Get(player.MOVE_RIGHT) && !ent.ActionBuffer.Get(player.MOVE_LEFT) {
		rec = pm[entity.KNIGHT].States["walking"].GetFrame("right", ent.FrameCount)
	} else {
		rec = pm[entity.KNIGHT].States["idling"].GetFrame("down", ent.FrameCount)
		walking = false
	}

	height := cam.ApplyOffsetF32(float32(rec.H / 2))
	width := cam.ApplyOffsetF32(float32(rec.W / 2))
	if walking {
		err := r.CopyF(pm[entity.KNIGHT].States["walking"].GetTexture(), &rec, &sdl.FRect{
			X: offseted.X,
			Y: offseted.Y,
			W: width,
			H: height,
		})
		if err != nil {
			lg.Error.Println("DROPING ERROR")
			return err
		}
	} else {
		err := r.CopyF(pm[entity.KNIGHT].States["idling"].GetTexture(), &rec, &sdl.FRect{
			X: offseted.X,
			Y: offseted.Y,
			W: width,
			H: height,
		})
		if err != nil {
			lg.Error.Println("DROPING ERROR")
			return err
		}
	}
	r.SetDrawColor(255, 255, 255, 255)
	r.DrawRectF(
		&sdl.FRect{
			X: offseted.X,
			Y: offseted.Y,
			W: width,
			H: height,
		},
	)
	ent.FrameCount += delta
	return nil
}
