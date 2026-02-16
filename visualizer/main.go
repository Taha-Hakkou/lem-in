package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// TODO:
// 	- move directions
//  - handling errors from lem-in
//  - example 2 & 3 not working properly !

func main() {
	// check whether stdin is coming from a pipe/file or from a terminal (TTY)
	stat, _ := os.Stdin.Stat()
	isPiped := (stat.Mode() & os.ModeCharDevice) == 0

	if !isPiped || len(os.Args) > 1 {
		fmt.Println("Usage: ./lem-in [FILE] | ./visualizer")
		return
	}

	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err) // why panic ?
	}
	lines := strings.Split(string(bytes), "\n")

	// need to check global error of lem-in

	parseData(lines)

	for _, p := range rooms {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	pos := make(map[string]Room)
	for id, p := range rooms {
		gx := (p.x - minX) * scale
		gy := (p.y - minY) * scale // to invert Y, use (maxY - p.y) instead
		pos[id] = Room{gx, gy}     // is id converted implicitly ???
	}
	// fmt.Println(pos)

	// ---------- CANVAS ----------
	for id, p := range pos {
		var w int = p.x + 2 + len(id) // 2 = opening & closing brackets
		if w > width {
			width = w
		}
		if p.y+1 > height {
			height = p.y + 1
		}
	}
	// fmt.Println(height, width)

	canvas := make([][]rune, height)
	for i := range canvas {
		canvas[i] = make([]rune, width)
		for j := range canvas[i] {
			canvas[i][j] = ' '
		}
	}
	// can i populate them in single line
	// + add space between lines so that slaches andd backslaches align properly !

	// ---------- DRAW ROOMS ----------
	for id, p := range pos {
		label := fmt.Sprintf("[%s]", id)
		for i, ch := range label {
			canvas[p.y][p.x+i] = ch
		}
	}

	// ---------- DRAW LINES ----------
	for _, e := range links {
		a, b := e[0], e[1]
		p1, p2 := pos[a], pos[b]
		drawLine(canvas, p1.x, p1.y, p2.x, p2.y)
	}

	flush(canvas)
	action = true
	Animate(canvas)
}
