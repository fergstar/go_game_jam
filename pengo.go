package main

import tl "github.com/JoelOtter/termloop"

var direction Direction
var collidedIceblock *Iceblock

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
			direction = RIGHT
			pengo.r.SetPosition(pengo.px+1, pengo.py)
			break
		case tl.KeyArrowLeft:
			direction = LEFT
			pengo.r.SetPosition(pengo.px-1, pengo.py)
			break
		case tl.KeyArrowUp:
			direction = UP
			pengo.r.SetPosition(pengo.px, pengo.py-1)
			break
		case tl.KeyArrowDown:
			direction = DOWN
			pengo.r.SetPosition(pengo.px, pengo.py+1)
			break
		case tl.KeySpace:

			var posX, posY = collidedIceblock.r.Position()
			switch direction {
			case LEFT:
				collidedIceblock.r.SetPosition(posX-1, posY)
				break
			case UP:
				collidedIceblock.r.SetPosition(posX, posY-1)
				break
			case RIGHT:
				collidedIceblock.r.SetPosition(posX+1, posY)
				break
			case DOWN:
				collidedIceblock.r.SetPosition(posX, posY+1)
				break
			}

			//  	if another iceblock next door
			// 			destroy iceblock next to pengo
			//  	else push iceblock in direction
			// else do nothing
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
	if iceblock, ok := collision.(*Iceblock); ok {
		collidedIceblock = iceblock
		pengo.r.SetPosition(pengo.px, pengo.py)

	}

}
