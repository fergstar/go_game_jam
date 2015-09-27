package main

type Direction int

const (
	NONE Direction = 1 + iota
	LEFT
	UP
	RIGHT
	DOWN
)

var directions = [...]string{
	"NONE",
	"LEFT",
	"UP",
	"RIGHT",
	"DOWN",
}

func (d Direction) String() string {
	return directions[d-1]
}
