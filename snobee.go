package main

import tl "github.com/JoelOtter/termloop"

type Snobee struct {
	r *tl.Rectangle
	x int
	y int
	g *tl.Game
}

func (s *Snobee) Size() (int, int) {
	return s.r.Size()
}

func (s *Snobee) Position() (int, int) {
	return s.r.Size()
}

func (s *Snobee) Draw(screen *tl.Screen) {
	s.r.Draw(screen)
}

func (s *Snobee) Tick(e tl.Event) {

}

func (s *Snobee) Collide(c tl.Physical) {

}
