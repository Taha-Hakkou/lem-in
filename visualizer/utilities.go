package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	currentRooms = make(map[int]string)
	m            sync.Mutex
)

func worker(canvas [][]rune, move [2]string, wg *sync.WaitGroup) {
	defer wg.Done() // Signal the WaitGroup when this goroutine finishes

	antId, _ := strconv.Atoi(move[0])
	m.Lock()
	roomname, ok := currentRooms[antId]
	currentRooms[antId] = move[1]
	m.Unlock()
	if !ok {
		roomname = start
	}
	room1 := rooms[roomname]
	room2 := rooms[move[1]]
	x1, y1 := (room1.x-minX)*scale, (room1.y-minY)*scale
	x2, y2 := (room2.x-minX)*scale, (room2.y-minY)*scale
	drawLine(canvas, x1, y1, x2, y2)
}

func Animate(canvas [][]rune) {
	for _, step := range steps {
		var wg sync.WaitGroup
		for _, move := range step {
			wg.Add(1)
			go worker(canvas, move, &wg)
		}
		wg.Wait()
		time.Sleep(2 * time.Second)
		reset()
		flush(canvas)
	}
}

func flush(canvas [][]rune) {
	// ---------- PRINT ----------
	for _, row := range canvas {
		line := string(row)
		line = strings.Replace(line, "["+start+"]", "\033[31m"+"["+start+"]"+"\033[0m", 1) // red ansi
		line = strings.Replace(line, "["+end+"]", "\033[32m"+"["+end+"]"+"\033[0m", 1)     // green ansi
		fmt.Println(line)
	}
}

func reset() {
	fmt.Printf("\033[%dA", height)
}

var changingCells = []string{}

func deplace(canvas [][]rune, x, y int) {
	for canvas[y][x] == '•' {
		time.Sleep(100 * time.Millisecond)
	}
	tmp := canvas[y][x]
	canvas[y][x] = '•' // '🐜'
	m.Lock()
	reset()
	flush(canvas)
	m.Unlock()
	time.Sleep(SLEEP)
	canvas[y][x] = tmp
}

// \r        → start of line
// \033[1A   → up 1 line
// \033[1B   → down 1 line
// \033[2K   → clear whole line
// \033[K    → clear from cursor right
