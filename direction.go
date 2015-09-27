package main

type Direction int

const (
	LEFT Direction = 1 + iota
	UP
	RIGHT
	DOWN
)

var directions = [...]string{
	"LEFT",
	"UP",
	"RIGHT",
	"DOWN",
}

func (d Direction) String() string {
	return directions[d-1]
}
