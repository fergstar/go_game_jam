package main

import tl "github.com/JoelOtter/termloop"

type Pengo struct {
	entity *tl.Entity
	prevX  int
	prevY  int
}

func (pengo *Pengo) Draw(screen *tl.Screen) {
	pengo.entity.Draw(screen)
}

func (pengo *Pengo) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		pengo.prevX, pengo.prevY = pengo.entity.Position()
		switch event.Key {
		case tl.KeyArrowRight:
			pengo.entity.SetPosition(pengo.prevX+1, pengo.prevY)
			break
		case tl.KeyArrowLeft:
			pengo.entity.SetPosition(pengo.prevX-1, pengo.prevY)
			break
		case tl.KeyArrowUp:
			pengo.entity.SetPosition(pengo.prevX, pengo.prevY-1)
			break
		case tl.KeyArrowDown:
			pengo.entity.SetPosition(pengo.prevX, pengo.prevY+1)
			break

		}
	}
}

func (pengo *Pengo) Size() (int, int) {
	return pengo.entity.Size()
}

func (pengo *Pengo) Position() (int, int) {
	return pengo.entity.Position()
}

func (pengo *Pengo) Collide(collision tl.Physical) {
	if _, ok := collision.(*tl.Rectangle); ok {
		pengo.entity.SetPosition(pengo.prevX, pengo.prevY)
	}
}
