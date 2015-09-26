package main

import tl "github.com/JoelOtter/termloop"

func NewPengo(x, y int, color tl.Attr, g *tl.Game, w, h, score int, scoretext *tl.Text) *Pengo {
	return &Pengo{
		r:         tl.NewRectangle(x, y, 1, 1, color),
		g:         g,
		w:         w,
		h:         h,
		score:     score,
		scoretext: scoretext,
	}

}

func BuildLevel(g *tl.Game, w, h, score int) {

	maze := generateMaze(w, h)
	l := tl.NewBaseLevel(tl.Cell{})
	l.SetOffset(30, 15)
	g.Screen().SetLevel(l)
	g.Log("Building level with width %d and height %d", w, h)

	for i, row := range maze {
		for j, path := range row {
			if path == '*' {
				// adds an iceblock but different to original pengo.
				// walls are not iceblocks.
				// TODO: change to use iceblock entity
				l.AddEntity(tl.NewRectangle(i, j, 1, 1, tl.ColorBlue))
			} else if path == 'P' {
				col := tl.RgbTo256Color(0xff, 0, 0)
				scoretext := tl.NewText(0, 1, "Pengo", tl.ColorWhite, tl.ColorBlack)
				g.Screen().AddEntity(scoretext)
				l.AddEntity(NewPengo(i, j, col, g, w, h, 0, scoretext))
			} //else if path == 'L' {
			//	l.AddEntity(tl.NewRectangle(i, j, 1, 1, tl.ColorBlue))
			//}
		}
	}
}

func main() {

	game := tl.NewGame()
	BuildLevel(game, 13, 15, 0)
	game.Start()

}
