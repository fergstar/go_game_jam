package main

import tl "github.com/JoelOtter/termloop"

type Iceblock struct {
	entity *tl.Entity
}

func (iceblock *Iceblock) Draw(screen *tl.Screen) {
	iceblock.entity.Draw(screen)
}
