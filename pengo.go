package main

import tl "github.com/JoelOtter/termloop"

type Pengo struct {
	r  *tl.Rectangle
	x  int // previous x position
	y  int // previous y position
	g  *tl.Game
	d  Direction
	ib *Iceblock
}

func (pengo *Pengo) Draw(screen *tl.Screen) {
	pengo.r.Draw(screen)
}

func (p *Pengo) Tick(event tl.Event) {
	if event.Type == tl.EventKey {

		// set previous position
		p.x, p.y = p.r.Position()

		// capture key events
		switch event.Key {
		case tl.KeyArrowRight:
			// set pengo's direction he is travelling.
			p.d = RIGHT
			// if pengo moves we want to unassign the iceblock that he
			// originally collided with.
			p.ib = nil
			// update pengo's position in the direction he is heading
			p.r.SetPosition(p.x+1, p.y)
			break
		case tl.KeyArrowLeft:
			p.d = LEFT
			p.ib = nil
			p.r.SetPosition(p.x-1, p.y)
			break
		case tl.KeyArrowUp:
			p.d = UP
			p.ib = nil
			p.r.SetPosition(p.x, p.y-1)
			break
		case tl.KeyArrowDown:
			p.d = DOWN
			p.ib = nil
			p.r.SetPosition(p.x, p.y+1)
			break
		case tl.KeySpace: // push the block

			// confirm that pengo has collided with a block
			if p.ib != nil {
				// if pengo has then set the iceblock to
				// travel in the direction pengo was going
				p.ib.direction = p.d
				// set the iceblock's current position as its start position
				p.ib.startPos.x, p.ib.startPos.y = p.ib.Position()
			}
			break
		}
	}
}

func (p *Pengo) Size() (int, int) {
	return p.r.Size()
}

func (p *Pengo) Position() (int, int) {
	return p.r.Position()
}

func (p *Pengo) Collide(collision tl.Physical) {

	// check if pengo has collided with an iceblock.
	if cib, ok := collision.(*Iceblock); ok {
		if cib.r.Color() == tl.ColorBlue {

			p.g.Log("pengo collided with iceblock")
			// assign the collided iceblock to pengo
			p.ib = cib
		}
		// set pengo's poistion to previous position
		p.r.SetPosition(p.x, p.y)
	}
}
