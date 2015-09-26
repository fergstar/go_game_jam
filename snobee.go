package main

import tl "github.com/JoelOtter/termloop"

type Snobee struct {
	entity *tl.Entity
}

func (snobee *Snobee) Draw(screen *tl.Screen) {
	snobee.entity.Draw(screen)
}
