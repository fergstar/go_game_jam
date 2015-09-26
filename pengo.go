package main

import tl "github.com/JoelOtter/termloop"

type Pengo struct {
	r         *tl.Rectangle
	px        int
	py        int
	move      bool
	g         *tl.Game
	w         int // Width of maze
	h         int // Hieght of maze
	score     int
	scoretext *tl.Text
}

func (pengo *Pengo) Draw(screen *tl.Screen) {
	pengo.r.Draw(screen)
}

func (pengo *Pengo) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		pengo.px, pengo.py = pengo.r.Position()
		switch event.Key {
		case tl.KeyArrowRight:
			pengo.r.SetPosition(pengo.px+1, pengo.py)
			break
		case tl.KeyArrowLeft:
			pengo.r.SetPosition(pengo.px-1, pengo.py)
			break
		case tl.KeyArrowUp:
			pengo.r.SetPosition(pengo.px, pengo.py-1)
			break
		case tl.KeyArrowDown:
			pengo.r.SetPosition(pengo.px, pengo.py+1)
			break

		}
	}
}

func (pengo *Pengo) Size() (int, int) {
	return pengo.r.Size()
}

func (pengo *Pengo) Position() (int, int) {
	return pengo.r.Position()
}

func (pengo *Pengo) Collide(collision tl.Physical) {
	if rectangle, ok := collision.(*tl.Rectangle); ok {
		if rectangle.Color() == tl.ColorBlue {
			// Collision with walls
			pengo.r.SetPosition(pengo.px, pengo.py)
		} else if rectangle.Color() == tl.ColorWhite {
			// Collision with end - new level!
			//b.w += 1
			//b.h += 1
			//b.score += 1
			//buildLevel(b.g, 13, 15, b.score)
		}
	}

}
