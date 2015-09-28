package main

import tl "github.com/JoelOtter/termloop"

type Pengo struct {
	r  *tl.Rectangle
	x  int
	y  int
	g  *tl.Game
	d  Direction
	ib *Iceblock
}

func (pengo *Pengo) Draw(screen *tl.Screen) {
	pengo.r.Draw(screen)
}

func (p *Pengo) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		p.x, p.y = p.r.Position()
		switch event.Key {
		case tl.KeyArrowRight:
			p.d = RIGHT
			p.ib = nil
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
		case tl.KeySpace:
			if p.ib != nil {
				p.ib.direction = p.d
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
	if cib, ok := collision.(*Iceblock); ok {
		p.g.Log("pengo collided with iceblock")
		p.ib = cib
		p.r.SetPosition(p.x, p.y)
	}
}
