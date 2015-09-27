package main

import tl "github.com/JoelOtter/termloop"

type Iceblock struct {
	r         *tl.Rectangle
	px        int
	py        int
	moving    bool
	direction Direction
}

func (iceblock *Iceblock) Draw(screen *tl.Screen) {

	if iceblock.moving && iceblock.direction != NONE {
		switch iceblock.direction {
		case RIGHT:
			iceblock.r.SetPosition(iceblock.px+1, iceblock.py)
			break
		case LEFT:
			iceblock.r.SetPosition(iceblock.px-1, iceblock.py)
			break
		case UP:
			iceblock.r.SetPosition(iceblock.px, iceblock.py-1)
			break
		case DOWN:
			iceblock.r.SetPosition(iceblock.px, iceblock.py+1)
			break
		}

	}
	iceblock.r.Draw(screen)

}

func (iceblock *Iceblock) Tick(event tl.Event) {

}

func (iceblock *Iceblock) Size() (int, int) {
	return iceblock.r.Size()
}

func (iceblock *Iceblock) Position() (int, int) {
	return iceblock.r.Position()
}

func (iceblock *Iceblock) Collide(collision tl.Physical) {
	if _, ok := collision.(*Iceblock); ok {
		// collide with another Iceblock
		iceblock.moving = false
		iceblock.r.SetPosition(iceblock.px, iceblock.py)
	}
}
