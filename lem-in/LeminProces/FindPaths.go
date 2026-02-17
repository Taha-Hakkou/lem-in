package lemin

// PathFinder chooses DFS for sparse graphs, BFS for dense graphs
func PathFinder(rooms []*Room, links []Link) [][]*Room {
	value := float64(len(links)) / float64(len(rooms))

	if value < 1.5 {
		return FindAllPathsDFS(rooms)
	}
	return FindDisjointPathsBFS(rooms)
}

// FindAllPathsDFS finds all possible paths using recursive DFS
func FindAllPathsDFS(rooms []*Room) [][]*Room {
	// Locate start and end rooms
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

	// Recursive DFS with backtracking
	var dfs func(*Room)
	dfs = func(r *Room) {
		visited[r] = true
		path = append(path, r)

		if r == end {
			// Save copy of complete path
			tmp := append([]*Room{}, path...)
			paths = append(paths, tmp)
		} else {
			// Explore unvisited neighbors
			for _, n := range r.Relations {
				if !visited[n] {
					dfs(n)
				}
			}
		}

		// Backtrack
		path = path[:len(path)-1]
		visited[r] = false
	}

	dfs(start)
	return paths
}

// FindDisjointPathsBFS finds node-disjoint paths using BFS
func FindDisjointPathsBFS(rooms []*Room) [][]*Room {
	// Locate start and end rooms
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
	used := map[*Room]bool{} // Nodes used in previous paths

	// Find paths iteratively
	for {
		queue := [][]*Room{{start}}
		visited := map[*Room]bool{start: true}
		var foundPath []*Room

		// BFS to find shortest available path
		for len(queue) > 0 && foundPath == nil {
			path := queue[0]
			queue = queue[1:]
			cur := path[len(path)-1]

			if cur == end {
				foundPath = path
				break
			}

			// Explore neighbors (skip used nodes)
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

		// Mark intermediate nodes as used
		for i := 1; i < len(foundPath)-1; i++ {
			used[foundPath[i]] = true
		}
	}

	return allPaths
}
