package main

import tl "github.com/JoelOtter/termloop"

func NewPengo(x, y int, game *tl.Game) *Pengo {
	return &Pengo{
		r: tl.NewRectangle(x, y, 1, 1, tl.ColorRed),
		x: x,
		y: y,
		g: game,
		d: NONE,
	}
}

func NewIceBlock(x, y int, game *tl.Game, color tl.Attr) *Iceblock {
	return &Iceblock{
		r:      tl.NewRectangle(x, y, 1, 1, color),
		g:      game,
		x:      x,
		y:      y,
		update: 0.05,
	}
}

func NewSnobee(x, y int, game *tl.Game) *Snobee {
	return &Snobee{
		r: tl.NewRectangle(x, y, 1, 1, tl.ColorYellow),
		x: x,
		y: y,
		g: game,
	}
}

func BuildLevel(g *tl.Game, w, h, score int) {

	maze := generateMaze(w, h)
	l := tl.NewBaseLevel(tl.Cell{})
	l.SetOffset(30, 15)
	g.Screen().SetLevel(l)
	g.Log("Building level with width %d and height %d", w, h)
	scoretext := tl.NewText(0, 1, "Pengo", tl.ColorWhite, tl.ColorBlack)
	g.Screen().AddEntity(scoretext)
	for i, row := range maze {
		for j, path := range row {
			if path == '*' {
				// it's an iceblock
				l.AddEntity(NewIceBlock(i, j, g, tl.ColorBlue))
			} else if path == 'P' {
				// it's Pengo
				l.AddEntity(NewPengo(i, j, g))
			} else if path == 'S' {
				// it's a Snobee
				l.AddEntity(NewSnobee(i, j, g))
			}
		}
	}
}

func main() {

	game := tl.NewGame()
	BuildLevel(game, 13, 15, 0)

	game.SetDebugOn(true)

	game.Start()

}
