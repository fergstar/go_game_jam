package main

import tl "github.com/JoelOtter/termloop"

type Iceblock struct {
	r         *tl.Rectangle
	g         *tl.Game
	update    float64
	time      float64
	direction Direction
	startPos  Point
	endPos    Point
}

type Point struct {
	x, y int
}

func (i *Iceblock) Draw(s *tl.Screen) {

	// if the iceblock has a direction
	if i.direction != NONE {
		// then move it in the direction specified
		i.time += s.TimeDelta()
		if i.time > i.update {
			switch i.direction {
			case RIGHT:
				i.r.SetPosition(i.endPos.x+1, i.endPos.y)
				break
			case LEFT:
				i.r.SetPosition(i.endPos.x-1, i.endPos.y)
				break
			case UP:
				i.r.SetPosition(i.endPos.x, i.endPos.y-1)
				break
			case DOWN:
				i.r.SetPosition(i.endPos.x, i.endPos.y+1)
			}
			// updated its previous position
			i.endPos.x, i.endPos.y = i.r.Position()
			i.time -= i.update
		}
	}
	i.r.Draw(s)
}

func (i *Iceblock) Tick(event tl.Event) {

}

func (i *Iceblock) Size() (int, int) {
	return i.r.Size()
}

func (i *Iceblock) Position() (int, int) {
	return i.r.Position()
}

func (i *Iceblock) Collide(collision tl.Physical) {

	// iceblock has collided with another iceblock
	if cib, ok := collision.(*Iceblock); ok {
		i.g.Log("iceblock collided with an iceblock.")

		// if collide iceblock is not moving and the moving iceblock is
		if cib.direction == NONE && i.direction != NONE {
			// then place it next to the iceblock it has collided with
			switch i.direction {
			case RIGHT:
				i.r.SetPosition(cib.endPos.x-1, cib.endPos.y)
				break
			case LEFT:
				i.r.SetPosition(cib.endPos.x+1, cib.endPos.y)
				break
			case UP:
				i.r.SetPosition(cib.endPos.x, cib.endPos.y+1)
				break
			case DOWN:
				i.r.SetPosition(cib.endPos.x, cib.endPos.y-1)
				break
			}

			// stop the iceblock from moving
			i.direction = NONE

			// update its previous position
			i.endPos.x, i.endPos.y = i.r.Position()

			// if the iceblock hasn't moved than we assume that
			// its next to another iceblock and we remove it
			// a.k.a pengo has crushed the iceblock
			if i.endPos.x == i.startPos.x && i.endPos.y == i.startPos.y {
				i.g.Screen().Level().RemoveEntity(i)
			}

		}
	}

	// iceblock has collided with a Snobee
	if _, ok := collision.(*Snobee); ok {

		// send the snobee moving in the same direction as iceblock
		// we may need to collect additional snobee's as we move
		// when it hits another iceblock or wall then crush snobee/s
	}
}
