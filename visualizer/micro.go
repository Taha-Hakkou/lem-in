package main

import (
	"time"
)

var SLEEP = 200 * time.Millisecond

func hyphens(canvas [][]rune, y, x1, x2 int) {
	var step int = 1
	if x2 < x1 {
		step = -1
	}
	for x := x1 + step; (x2-x)*step > 0; x += step {
		// check because room name takes more than a cell
		if action {
			deplace(canvas, x, y)
		} else if canvas[y][x] == ' ' {
			canvas[y][x] = '-'
		}
	}
}

func pipes(canvas [][]rune, x, y1, y2 int) int {
	var step int = 1
	var y int
	if y2 < y1 {
		step = -1
	}
	for y = y1 + step; (y2-y)*step > 0; y += step {
		// moved 1 cell to the right to be appropriately aligned
		if action {
			deplace(canvas, x, y)
		} else if canvas[y][x] == ' ' || canvas[y][x] == '-' || canvas[y][x] == '_' || canvas[y][x] == '/' || canvas[y][x] == '\\' {
			canvas[y][x] = '|'
		}
	}
	return y
}

func underscores(canvas [][]rune, y, x1, x2 int) int {
	// same as hyphens
	var step int = 1
	var x int
	if x2 < x1 {
		step = -1
	}
	for x = x1 + step; (x2-x)*step > 0; x += step {
		// check because room name takes more than a cell
		if action {
			deplace(canvas, x, y)
		} else if canvas[y][x] == ' ' || canvas[y][x] == '-' {
			canvas[y][x] = '_'
		}
	}
	return x
}

func backslashes(canvas [][]rune, x, y, x2, y2 int) (int, int) {
	var step int = 1
	if x > x2 {
		step = -1
	}
	// if x2 > x : y <= y2 && x <= end
	// else      : y > y2 && x > end
	for (y2-y)*step > 0 && (x2-x)*step > 0 { // or equal (if step = 1): not implemented
		if action {
			deplace(canvas, x, y)
		} else if canvas[y][x] == ' ' || canvas[y][x] == '_' || canvas[y][x] == '-' {
			canvas[y][x] = '\\'
		}
		x += step
		y += step
	}
	if step == 1 {
		if action {
			deplace(canvas, x, y)
		} else if canvas[y][x] == ' ' || canvas[y][x] == '_' || canvas[y][x] == '-' {
			canvas[y][x] = '\\'
		}
		x++
	}
	return x, y
}

func slashes(canvas [][]rune, x, y, x2, y2 int) (int, int) {
	var step int = 1
	if x > x2 {
		step = -1
	}
	// if x2 > x : y > y2 && x < end
	// else      : y <= y2 && x >= end
	for (x2-x)*step > 0 && (y2-y)*step < 0 { // or equal (if step = -1): not implemented
		if action {
			deplace(canvas, x, y)
		} else if canvas[y][x] == ' ' || canvas[y][x] == '_' || canvas[y][x] == '-' {
			canvas[y][x] = '/'
		}
		x += step
		y -= step
	}
	if step == -1 {
		if action {
			deplace(canvas, x, y)
		} else if canvas[y][x] == ' ' || canvas[y][x] == '_' || canvas[y][x] == '-' {
			canvas[y][x] = '/'
		}
		x--
	}
	return x, y
}
