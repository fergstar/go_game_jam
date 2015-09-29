package main

import tl "github.com/JoelOtter/termloop"

// TODO move to pengo.go
func NewPengo(x, y int, game *tl.Game) *Pengo {
	return &Pengo{
		r: tl.NewRectangle(x, y, 1, 1, tl.ColorRed),
		x: x,
		y: y,
		g: game,
		d: NONE,
	}
}

// TODO move to iceblock.go
func NewIceBlock(x, y int, game *tl.Game, color tl.Attr) *Iceblock {
	return &Iceblock{
		r:      tl.NewRectangle(x, y, 1, 1, color),
		g:      game,
		x:      x,
		y:      y,
		update: 0.05,
	}
}

// TODO move to snobee.go
func NewSnobee(x, y int, game *tl.Game) *Snobee {
	return &Snobee{
		r: tl.NewRectangle(x, y, 1, 1, tl.ColorYellow),
		x: x,
		y: y,
		g: game,
	}
}

// TODO need to be able to build multiple levels
// maybe look at adding a 'R' key during debug to reload/refresh level
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
			// 'R' it's a Diamond iceblock
			// 's' it's a snobee egg iceblock
		}
	}
}

func main() {

	game := tl.NewGame()

	// pengo default maze size is 13x15
	BuildLevel(game, 13, 15, 0)

	game.SetDebugOn(false)

	game.Start()

}
