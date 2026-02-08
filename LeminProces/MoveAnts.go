package lemin

import (
	"fmt"
	"strings"
)

func MoveAnts(paths [][]*Room, nAnts int) {
	if len(paths) == 0 || nAnts == 0 {
		return
	}

	// sort paths by length
	for i := 0; i < len(paths)-1; i++ {
		for j := 0; j < len(paths)-i-1; j++ {
			if len(paths[j]) > len(paths[j+1]) {
				paths[j], paths[j+1] = paths[j+1], paths[j]
			}
		}
	}

	//  choose best disjoint set
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
		time := 0
		for k := range set {
			if dist[k] > 0 {
				t := len(set[k]) - 1 + dist[k] - 1
				if t > time {
					time = t
				}
			}
		}

		if bestTime == -1 || time < bestTime {
			bestTime = time
			best = set
		}
	}

	if len(best) == 0 {
		return
	}

	//  assign ants
	dist := distribute(best, nAnts)
	var ants []Ant
	id := 1

	for i := range best {
		for c := 0; c < dist[i]; c++ {
			ants = append(ants, Ant{
				Number: id,
				Pos:    0,
				Path:   best[i],
			})
			id++
		}
	}

	//  simulate turns
	finished := 0

	for finished < nAnts {
		used := make(map[string]bool)
		var moves []string

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
			moves = append(moves,
				fmt.Sprintf("L%d-%s", a.Number, next.Name),
			)

			if next.Role != "end" {
				used[next.Name] = true
			}

			if a.Pos == len(a.Path)-1 {
				finished++
			}
		}

		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
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
