package lemin

import (
	"fmt"
	"strings"
)

func PathFinder(rooms []*Room, links []Link) [][]*Room {
	avg := float64(len(links)*2) / float64(len(rooms))
	if avg < 2.5 {
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
		}
		if r.Role == "end" {
			end = r
		}
	}
	if start == nil || end == nil {
		return nil
	}

	var all [][]*Room
	used := map[*Room]bool{}

	for {
		p := bfsFindPath(start, end, used)
		if p == nil {
			break
		}
		all = append(all, p)
		for i := 1; i < len(p)-1; i++ {
			used[p[i]] = true
		}
	}

	return all
}

func bfsFindPath(start, end *Room, used map[*Room]bool) []*Room {
	queue := [][]*Room{{start}}
	visited := map[*Room]bool{start: true}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		cur := p[len(p)-1]

		if cur == end {
			return p
		}

		for _, n := range cur.Relations {
			if visited[n] || (used[n] && n != end) {
				continue
			}
			visited[n] = true
			np := append(append([]*Room{}, p...), n)
			queue = append(queue, np)
		}
	}

	return nil
}

func MoveAnts(paths [][]*Room, ants int) {
	if len(paths) == 0 || ants == 0 {
		return
	}

	sortPaths(paths)
	best := selectBest(paths, ants)
	if len(best) == 0 {
		return
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

func sortPaths(paths [][]*Room) {
	for i := 0; i < len(paths)-1; i++ {
		for j := 0; j < len(paths)-i-1; j++ {
			if len(paths[j]) > len(paths[j+1]) {
				paths[j], paths[j+1] = paths[j+1], paths[j]
			}
		}
	}
}

func selectBest(paths [][]*Room, ants int) [][]*Room {
	var best [][]*Room
	bestTime := -1

	for i := range paths {
		set := [][]*Room{paths[i]}
		for j := range paths {
			if i != j && disjoint(set, paths[j]) {
				set = append(set, paths[j])
			}
		}

		dist := distribute(set, ants)
		t := calcTime(set, dist)

		if bestTime == -1 || t < bestTime {
			bestTime = t
			best = set
		}
	}

	return best
}

func calcTime(paths [][]*Room, dist []int) int {
	max := 0
	for i := range paths {
		if dist[i] > 0 {
			t := len(paths[i]) - 1 + dist[i] - 1
			if t > max {
				max = t
			}
		}
	}
	return max
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
