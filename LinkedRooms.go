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

func MoveAnts(paths [][]*Room, nAnts int) {
	if len(paths) == 0 || nAnts == 0 {
		return
	}

	// sort by length
	for i := 0; i < len(paths)-1; i++ {
		for j := 0; j < len(paths)-i-1; j++ {
			if len(paths[j]) > len(paths[j+1]) {
				paths[j], paths[j+1] = paths[j+1], paths[j]
			}
		}
	}

	// build best disjoint set
	var best [][]*Room
	bestTime := -1

	for i := range paths {
		set := [][]*Room{paths[i]}
		for j := range paths {
			if i != j && disjoint(set, paths[j]) {
				set = append(set, paths[j])
			}
		}

		dist := distribute(set, nAnts)
		t := maxTime(set, dist)

		if bestTime == -1 || t < bestTime {
			bestTime = t
			best = set
		}
	}

	if len(best) == 0 {
		return
	}

	// distribute ants SAFELY
	dist := distribute(best, nAnts)

	var ants []Ant
	id := 1
	for i := 0; i < len(best); i++ {
		for c := 0; c < dist[i]; c++ {
			ants = append(ants, Ant{
				Number: id,
				Pos:    0,
				Path:   best[i],
			})
			id++
		}
	}

	// simulate
	finished := 0
	for finished < nAnts {
		used := make(map[string]bool)
		var line []string

		for i := range ants {
			a := &ants[i]

			if a.Pos == len(a.Path)-1 {
				continue
			}

			next := a.Path[a.Pos+1]
			if next.Role != "end" && used[next.Name] {
				continue
			}

			a.Pos++
			line = append(line,
				fmt.Sprintf("L%d-%s", a.Number, next.Name))

			if next.Role != "end" {
				used[next.Name] = true
			}

			if a.Pos == len(a.Path)-1 {
				finished++
			}
		}

		if len(line) > 0 {
			fmt.Println(strings.Join(line, " "))
		}
	}
}


func disjoint(set [][]*Room, p []*Room) bool {
	for _, s := range set {
		for i := 1; i < len(s)-1; i++ {
			for j := 1; j < len(p)-1; j++ {
				if s[i].Name == p[j].Name {
					return false
				}
			}
		}
	}
	return true
}

func distribute(paths [][]*Room, ants int) []int {
	d := make([]int, len(paths))
	for ants > 0 {
		best := 0
		for i := range paths {
			if len(paths[i])+d[i] < len(paths[best])+d[best] {
				best = i
			}
		}
		d[best]++
		ants--
	}
	return d
}

func maxTime(paths [][]*Room, d []int) int {
	m := 0
	for i := range paths {
		if d[i] > 0 {
			t := len(paths[i]) - 1 + d[i] - 1
			if t > m {
				m = t
			}
		}
	}
	return m
}
