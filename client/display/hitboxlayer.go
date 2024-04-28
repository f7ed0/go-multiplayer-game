package display

import (
	hitbox "github.com/f7ed0/go-multiplayer-game/commons/Hitbox"
)

func (w *Window) hiboxLayer() {
	w.renderer.SetDrawColor(255, 0, 0, 255)
	for _, poly := range w.GameMap.Walls {
		w.DrawHitboxes(poly)
	}
	w.renderer.SetDrawColor(0, 0, 255, 255)
	for _, poly := range w.GameMap.Holes {
		w.DrawHitboxes(poly)
	}
	for _, hits := range w.Me.HitBoxes {
		w.DrawHitboxes(hitbox.Hitbox{hits.Polygon.OffsetPolygon(w.Me.Position)})
	}

}

func (w *Window) DrawHitboxes(p hitbox.Hitbox) {
	for i := 0; i < len(p.Points)-1; i++ {
		pi := w.applyOffset(p.Points[i])
		pi1 := w.applyOffset(p.Points[i+1])
		//lg.Verbose.Println(p.Points[i], offset, pi)
		w.renderer.DrawLineF(pi.X, pi.Y, pi1.X, pi1.Y)
	}
	pl := w.applyOffset(p.Points[len(p.Points)-1])
	pf := w.applyOffset(p.Points[0])
	w.renderer.DrawLineF(pl.X, pl.Y, pf.X, pf.Y)
}
