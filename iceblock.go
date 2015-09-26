package main

import tl "github.com/JoelOtter/termloop"

type Iceblock struct {
	r *tl.Rectangle
}

func (iceblock *Iceblock) Draw(screen *tl.Screen) {
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
}
