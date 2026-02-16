package main

func drawLine(canvas [][]rune, x1, y1, x2, y2 int) {
	if y1 == y2 {
		hyphens(canvas, y1, x1, x2)
		return
	}
	if x1 == x2 {
		pipes(canvas, x1+1, y1, y2)
		return
	}

	var x, y int = x1, y1
	var startWithPipes bool

	// ------------------------------------
	// if y1+y2 < height && x1+x2 < width {
	// 	startWithPipes = true
	// }
	if y1+y2 >= height {
		startWithPipes = true
	}
	if y2 < y1 {
		startWithPipes = !startWithPipes
	}
	// ------------------------------------

	// 1
	if startWithPipes {
		// pipes
		x++ // to be aligned approprietaly with room name
		if x2 > x {
			if y2 > y {
				if y2-y >= x2-x {
					y = pipes(canvas, x, y, y2-x2+x+1) // no need to do y = y1
				}
			} else {
				if y-y2 >= x2-x {
					y = pipes(canvas, x, y, y2+x2-x-1)
				}
			}
			if y < y1-1 || y > y1+1 { // check if the loop was entered
				x++
			}
		} else {
			if y2 > y {
				if y2-y >= x-x2 {
					y = pipes(canvas, x, y, y2+x2-x+1) // no need to do y = y1
				}
			} else {
				if y-y2 >= x-x2 {
					y = pipes(canvas, x, y, y2-x2+x-1)
				}
			}
			if y < y1-1 || y > y1+1 { // check if the loop was entered
				x--
			}
		}
	} else {
		// underscores
		if y2 > y {
			if x2 > x {
				if x2-x >= y2-y {
					x = underscores(canvas, y, x, x2-y2+y+1) // no need to do y = y1
				}
			} else {
				if x-x2 >= y2-y {
					x = underscores(canvas, y, x, x2+y2-y-1)
				}
			}
			if x < x1-1 || x > x1+1 { // only if going down after underscores.
				y++
			}
		} else {
			if x2 > x {
				if x2-x >= y-y2 {
					x = underscores(canvas, y, x, x2+y2-y+1) // no need to do y = y1
				}
			} else {
				if x-x2 >= y-y2 {
					x = underscores(canvas, y, x, x2-y2+y-1)
				}
			}
			// if x < x1-1 || x > x1+1 { // check if the loop was entered
			// 	y--
			// }
		}
	}

	// 2: slashes & backslashes
	var end int = x2     // for pipe alignement
	if !startWithPipes { // end with pipes
		end++
	}
	if (x2-x)*(y2-y) > 0 {
		// backslashes
		x, y = backslashes(canvas, x, y, end, y2)
		// if startWithPipes {
		// 	y++
		// }
	} else {
		// slashes
		x, y = slashes(canvas, x, y, end, y2) // end necessary ?
		// if startWithPipes {
		// 	y++
		// }
	}

	// 3
	if startWithPipes {
		// underscores
		if x2 > x {
			underscores(canvas, y, x-1, x2)
		} else {
			underscores(canvas, y, x+1, x2)
		}
	} else {
		// pipes
		if y2 > y {
			pipes(canvas, x, y-1, y2)
		} else {
			pipes(canvas, x, y+1, y2)
		}
	}
}
