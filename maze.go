package main

import (
	"math/rand"
	"time"
)

// Maze generation taken from _examples/pyramid.go & modified
// needs more work to simulate maze from Pengo

type MazePoint struct {
	x int
	y int
	p *MazePoint
}

func (p *MazePoint) Opposite() *MazePoint {
	if p.x != p.p.x {
		return &MazePoint{x: p.x + (p.x - p.p.x), y: p.y, p: p}
	}
	if p.y != p.p.y {
		return &MazePoint{x: p.x, y: p.y + (p.y - p.p.y), p: p}
	}
	return nil
}

func adjacents(point *MazePoint, maze [][]rune) []*MazePoint {
	res := make([]*MazePoint, 0)
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if (i == 0 && j == 0) || (i != 0 && j != 0) {
				continue
			}
			if !isInMaze(point.x+i, point.y+j, len(maze), len(maze[0])) {
				continue
			}
			if maze[point.x+i][point.y+j] == '*' {
				res = append(res, &MazePoint{point.x + i, point.y + j, point})
			}
		}
	}
	return res
}

func isInMaze(x, y int, w, h int) bool {
	return x >= 0 && x < w &&
		y >= 0 && y < h
}

// Generates a maze using Prim's Algorithm
// https://en.wikipedia.org/wiki/Maze_generation_algorithm#Randomized_Prim.27s_algorithm
func generateMaze(w, h int) [][]rune {
	maze := make([][]rune, w)
	for row := range maze {
		maze[row] = make([]rune, h)
		for ch := range maze[row] {
			maze[row][ch] = '*'
		}
	}
	rand.Seed(time.Now().UnixNano())
	// Pengo always starts at 7,7
	point := &MazePoint{x: 7, y: 7}
	maze[point.x][point.y] = 'P'

	var last *MazePoint
	walls := adjacents(point, maze)
	for len(walls) > 0 {
		rand.Seed(time.Now().UnixNano())
		wall := walls[rand.Intn(len(walls))]
		for i, w := range walls {
			if w.x == wall.x && w.y == wall.y {
				walls = append(walls[:i], walls[i+1:]...)
				break
			}
		}
		opp := wall.Opposite()
		if isInMaze(opp.x, opp.y, w, h) && maze[opp.x][opp.y] == '*' {
			maze[wall.x][wall.y] = '.'
			maze[opp.x][opp.y] = '.'
			walls = append(walls, adjacents(opp, maze)...)
			last = opp
		}
	}
	maze[last.x][last.y] = 'S'
	bordered := make([][]rune, len(maze)+2)
	for r := range bordered {
		bordered[r] = make([]rune, len(maze[0])+2)
		for c := range bordered[r] {
			if r == 0 || r == len(maze)+1 || c == 0 || c == len(maze[0])+1 {
				bordered[r][c] = '*'
			} else {
				bordered[r][c] = maze[r-1][c-1]
			}
		}
	}
	return bordered
}
