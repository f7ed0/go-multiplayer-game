package drawableentity

import (
	"math"

	"github.com/f7ed0/go-multiplayer-game/client/handleplayer"
	"github.com/f7ed0/go-multiplayer-game/client/handleplayer/camera"
	"github.com/f7ed0/go-multiplayer-game/commons/entity"
	"github.com/f7ed0/go-multiplayer-game/commons/entity/player"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
	"github.com/f7ed0/golog/lg"
	"github.com/veandco/go-sdl2/sdl"
)

var FrameCounter int = 0

func Draw(ent *handleplayer.DispPlayer, r *sdl.Renderer, cam camera.Camera, pm map[entity.EntityType]EntityModel, delta int) error {
	boundaries := r.GetViewport()

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
		walking = false
		if ent.Orientation > math.Pi/4 && ent.Orientation < 3*math.Pi/4 {
			rec = pm[entity.KNIGHT].States["idling"].GetFrame("down", ent.FrameCount)
		} else if ent.Orientation < -math.Pi/4 && ent.Orientation > -3*math.Pi/4 {
			rec = pm[entity.KNIGHT].States["idling"].GetFrame("up", ent.FrameCount)
		} else if ent.Orientation > -math.Pi/4 && ent.Orientation < math.Pi/4 {
			rec = pm[entity.KNIGHT].States["idling"].GetFrame("right", ent.FrameCount)
		} else {
			rec = pm[entity.KNIGHT].States["idling"].GetFrame("left", ent.FrameCount)
		}

	}

	offseted := cam.ApplyOffset(objects.Point{X: ent.Position.X - float32(rec.W/4), Y: ent.Position.Y}, int(boundaries.W), int(boundaries.H))

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
	ent.FrameCount += delta
	return nil
}
