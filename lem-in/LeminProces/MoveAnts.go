package lemin

import (
	"fmt"
	"strings"
)

func PathFinder(rooms []*Room, links []Link) [][]*Room {
	ratio := float64(len(links)) / float64(len(rooms))

	// Si peu de liens par rapport aux salles → DFS
	// Si beaucoup de liens → BFS
	if ratio < 1.5 {
		
		return FindAllPathsDFS(rooms)
	}
	
	
	return FindDisjointPathsBFS(rooms)
}

func FindAllPathsDFS(rooms []*Room) [][]*Room {
	var start, end *Room
	for _, r := range rooms {
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
	visited := map[*Room]bool{}

	var dfs func(*Room)
	dfs = func(r *Room) {
		visited[r] = true
		path = append(path, r)

		if r == end {
			tmp := append([]*Room{}, path...)
			paths = append(paths, tmp)
		} else {
			for _, n := range r.Relations {
				if !visited[n] {
					dfs(n)
				}
			}
		}

		path = path[:len(path)-1]
		visited[r] = false
	}

	dfs(start)
	return paths
}

func FindDisjointPathsBFS(rooms []*Room) [][]*Room {
	var start, end *Room
	for _, r := range rooms {
		if r.Role == "start" {
			start = r
		} else if r.Role == "end" {
			end = r
		}
	}
	if start == nil || end == nil {
		return nil
	}

	var allPaths [][]*Room
	used := map[*Room]bool{}

	for {
		queue := [][]*Room{{start}}
		visited := map[*Room]bool{start: true}
		var foundPath []*Room

		for len(queue) > 0 && foundPath == nil {
			path := queue[0]
			queue = queue[1:]
			cur := path[len(path)-1]

			if cur == end {
				foundPath = path
				break
			}

			for _, next := range cur.Relations {
				if visited[next] || (used[next] && next != end) {
					continue
				}
				visited[next] = true
				newPath := append([]*Room{}, path...)
				newPath = append(newPath, next)
				queue = append(queue, newPath)
			}
		}

		if foundPath == nil {
			break
		}

		allPaths = append(allPaths, foundPath)
		for i := 1; i < len(foundPath)-1; i++ {
			used[foundPath[i]] = true
		}
	}

	return allPaths
}
func MoveAnts(paths [][]*Room, ants int) {
	if len(paths) == 0 || ants == 0 {
		return
	}

	for i := 0; i < len(paths)-1; i++ {
		for j := 0; j < len(paths)-i-1; j++ {
			if len(paths[j]) > len(paths[j+1]) {
				paths[j], paths[j+1] = paths[j+1], paths[j]
			}
		}
	}

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

	dist := distribute(best, ants)

	var list []Ant
	id := 1
	for i := range best {
		for j := 0; j < dist[i]; j++ {
			list = append(list, Ant{id, 0, best[i]})
			id++
		}
	}

	done := 0
	for done < ants {
		used := map[string]bool{}
		var moves []string

		for i := range list {
			a := &list[i]

			if a.Pos == len(a.Path)-1 {
				continue
			}

			n := a.Path[a.Pos+1]
			if n.Role != "end" && used[n.Name] {
				continue
			}

			a.Pos++
			moves = append(moves, fmt.Sprintf("L%d-%s", a.Number, n.Name))

			if n.Role != "end" {
				used[n.Name] = true
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
