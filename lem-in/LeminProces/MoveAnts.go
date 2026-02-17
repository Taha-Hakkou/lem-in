package lemin

import (
	"fmt"
	"strings"
)

// MoveAnts simulates ant movement respecting room and tunnel rules
func MoveAnts(paths [][]*Room, ants int) {
	if len(paths) == 0 || ants == 0 {
		return
	}

	// Distribute ants across paths
	dist := distribute(paths, ants)

	// Create ant list with assigned paths
	var list []Ant
	id := 1
	for i := range paths {
		for j := 0; j < dist[i]; j++ {
			list = append(list, Ant{
				Number: id,
				Pos:    0,
				Path:   paths[i],
			})
			id++
		}
	}

	done := 0

	// Simulate turns
	for done < ants {
		usedRooms := map[string]bool{}
		usedTunnels := map[string]bool{}
		var moves []string

		// Try to move each ant
		for i := range list {
			a := &list[i]

			// Skip ants at end
			if a.Pos == len(a.Path)-1 {
				continue
			}

			current := a.Path[a.Pos]
			next := a.Path[a.Pos+1]
			tunnel := current.Name + "-" + next.Name

			// Check tunnel and room availability
			if usedTunnels[tunnel] {
				continue
			}
			if next.Role != "end" && usedRooms[next.Name] {
				continue
			}

			// Move ant
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

// disjoint checks if path shares intermediate nodes with set
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

// distribute assigns ants to minimize completion time
func distribute(paths [][]*Room, ants int) []int {
	d := make([]int, len(paths))

	// Greedy: assign to path with minimum (length + assigned)
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
