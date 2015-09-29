package main

import tl "github.com/JoelOtter/termloop"

type Iceblock struct {
	r         *tl.Rectangle
	g         *tl.Game
	x         int
	y         int
	update    float64
	time      float64
	direction Direction
}

func (i *Iceblock) Draw(s *tl.Screen) {

	// if the iceblock has a direction
	if i.direction != NONE {
		// then move it in the direction specified
		i.time += s.TimeDelta()
		if i.time > i.update {
			switch i.direction {
			case RIGHT:
				i.r.SetPosition(i.x+1, i.y)
				break
			case LEFT:
				i.r.SetPosition(i.x-1, i.y)
				break
			case UP:
				i.r.SetPosition(i.x, i.y-1)
				break
			case DOWN:
				i.r.SetPosition(i.x, i.y+1)
			}
			// updated its previous position
			i.x, i.y = i.r.Position()
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

		// TODO crush iceblock if its next to another iceblock
		// we can do this by setting its start position and its
		// end position (previous position)
		// if the start pos equals the end pos then the iceblock
		// hasn't moved and we can crush the iceblock.

		// if iceblock has direction
		if i.direction != NONE {
			// then place it next to the iceblock it has collided with
			switch i.direction {
			case RIGHT:
				i.r.SetPosition(cib.x-1, cib.y)
				break
			case LEFT:
				i.r.SetPosition(cib.x+1, cib.y)
				break
			case UP:
				i.r.SetPosition(cib.x, cib.y+1)
				break
			case DOWN:
				i.r.SetPosition(cib.x, cib.y-1)
				break
			}

			// stop the iceblock from moving
			i.direction = NONE

			// and update its previous position
			i.x, i.y = i.r.Position()

		}
	}

	// iceblock has collided with a Snobee
	if _, ok := collision.(*Snobee); ok {

		// send the snobee moving in the same direction as iceblock
		// we may need to collect additional snobee's as we move
		// when it hits another iceblock or wall then crush snobee/s
	}
}
