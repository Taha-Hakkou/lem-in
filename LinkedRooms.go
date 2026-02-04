package main

import (
	"fmt"
	"strings"
)

func GetRelatedRooms(Rooms []*Room, Links []Link) {
	for i := range Rooms {
		for _, l := range Links {
			if Rooms[i].Name == l.R1.Name {
				Rooms[i].Relations = append(Rooms[i].Relations, l.R2)
			}
			if Rooms[i].Name == l.R2.Name {
				Rooms[i].Relations = append(Rooms[i].Relations, l.R1)
			}
		}
	}
}

func FindAllPaths(Rooms []*Room) [][]*Room {
	var start, end *Room

	for _, r := range Rooms {
		if r.Role == "start" {
			start = r
		}
		if r.Role == "end" {
			end = r
		}
	}

	if start == nil || end == nil {
		return nil
	}

	var paths [][]*Room
	var path []*Room
	visited := make(map[*Room]bool)

	var DFS func(r *Room)
	DFS = func(r *Room) {
		visited[r] = true
		path = append(path, r)

		if r == end {
			tmp := make([]*Room, len(path))
			copy(tmp, path)
			paths = append(paths, tmp)
		} else {
			for _, n := range r.Relations {
				if !visited[n] {
					DFS(n)
				}
			}
		}

		// backtracking
		path = path[:len(path)-1]
		visited[r] = false
	}

	DFS(start)
	return paths
}

func MoveAnts(Paths [][]*Room, NAnts int) {
	for i := 0; i < len(Paths)-1; i++ {
		for j := 0; j < len(Paths)-i-1; j++ {
			if len(Paths[j]) > len(Paths[j+1]) {
				Paths[j], Paths[j+1] = Paths[j+1], Paths[j]
			}
		}
	}

	var Ants []Ant
	j := 0
	for i := 0; i < NAnts; i++ {
		if j == len(Paths) {
			j = 0
		}
		a := Ant{Number: i, Pos: 0, Path: Paths[j]}
		Ants = append(Ants, a)
		j++
	}
	finished := 0

	for finished < NAnts {
		var line []string

		for i := range Ants {
			ant := &Ants[i]

			if ant.Pos >= len(ant.Path)-1 {
				continue
			}

			nextRoom := ant.Path[ant.Pos+1]

			ant.Pos++
			line = append(line,
				fmt.Sprintf("L%d-%s", ant.Number, nextRoom.Name),
			)

			if nextRoom.Role == "end" {
				finished++
			}
		}

		if len(line) > 0 {
			fmt.Println(strings.Join(line, " "))
		}
	}
}
