package lemin

import (
	"fmt"
	"strings"
)

func MoveAnts(paths [][]*Room, ants int) {
	if len(paths) == 0 || ants == 0 {
		return
	}

	// Manual sort (shortest first)
	for i := 0; i < len(paths)-1; i++ {
		for j := 0; j < len(paths)-i-1; j++ {
			if len(paths[j]) > len(paths[j+1]) {
				paths[j], paths[j+1] = paths[j+1], paths[j]
			}
		}
	}

	// Build best disjoint set (max number of paths)
	var best [][]*Room
	for i := range paths {
		set := [][]*Room{paths[i]}
		for j := range paths {
			if i != j && disjoint(set, paths[j]) {
				set = append(set, paths[j])
			}
		}
		if len(set) > len(best) {
			best = set
		}
	}

	// Distribute ants
	dist := distribute(best, ants)

	// Create ants
	var list []Ant
	id := 1
	for i := range best {
		for j := 0; j < dist[i]; j++ {
			list = append(list, Ant{
				Number: id,
				Pos:    0,
				Path:   best[i],
			})
			id++
		}
	}

	done := 0

	for done < ants {
		usedRooms := map[string]bool{}
		usedTunnels := map[string]bool{}
		var moves []string

		for i := range list {
			a := &list[i]

			if a.Pos == len(a.Path)-1 {
				continue
			}

			current := a.Path[a.Pos]
			next := a.Path[a.Pos+1]
			tunnel := current.Name + "-" + next.Name

			if usedTunnels[tunnel] {
				continue
			}

			if next.Role != "end" && usedRooms[next.Name] {
				continue
			}

			a.Pos++
			moves = append(moves, fmt.Sprintf("L%d-%s", a.Number, next.Name))

			usedTunnels[tunnel] = true
			if next.Role != "end" {
				usedRooms[next.Name] = true
			}

			if a.Pos == len(a.Path)-1 {
				done++
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
		for i := 1; i < len(paths); i++ {
			if len(paths[i])+d[i] < len(paths[best])+d[best] {
				best = i
			}
		}
		d[best]++
		ants--
	}

	return d
}
