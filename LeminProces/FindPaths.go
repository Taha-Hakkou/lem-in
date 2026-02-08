package lemin

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


