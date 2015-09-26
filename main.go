package main

import tl "github.com/JoelOtter/termloop"

func main() {

	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorBlack,
	})

	level.AddEntity(tl.NewRectangle(10, 10, 13, 1, tl.ColorBlue))
	level.AddEntity(tl.NewRectangle(9, 10, 1, 17, tl.ColorBlue))
	level.AddEntity(tl.NewRectangle(23, 10, 1, 17, tl.ColorBlue))
	level.AddEntity(tl.NewRectangle(10, 26, 13, 1, tl.ColorBlue))

	pengo := Pengo{
		entity: tl.NewEntity(11, 11, 1, 1),
	}
	// Set the character at position (0, 0) on the entity.
	pengo.entity.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: 'P'})
	level.AddEntity(&pengo)

	game.Screen().SetLevel(level)

	game.Start()

}
