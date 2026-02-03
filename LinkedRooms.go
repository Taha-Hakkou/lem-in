package main

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

func FindPath(Rooms []*Room) bool {
	var Queue []*Room
	var start *Room

	for i := range Rooms {
		if Rooms[i].Role == "start" {
			start = Rooms[i]
			start.IsVisited = true
			start.Steps = 0
			break
		}
	}

	if start == nil {
		return false
	}

	Queue = append(Queue, start)
	pathFound := false

	for len(Queue) > 0 {
		current := Queue[0]
		Queue = Queue[1:]

		for _, neighbor := range current.Relations {
			if !neighbor.IsVisited {
				neighbor.IsVisited = true
				neighbor.Parent = append(neighbor.Parent, current)
				neighbor.Steps = current.Steps + 1
				Queue = append(Queue, neighbor)

				if neighbor.Role == "end" {
					pathFound = true
				}
			} else if neighbor.Steps == current.Steps+1 {
				alreadyHas := false
				for _, p := range neighbor.Parent {
					if p.Name == current.Name {
						alreadyHas = true
						break
					}
				}
				if !alreadyHas {
					neighbor.Parent = append(neighbor.Parent, current)
				}
			}
		}
	}

	return pathFound
}

func ExtractAllPaths(Rooms []*Room) [][]*Room {
	var end *Room

	// Trouver la salle de fin
	for i := range Rooms {
		if Rooms[i].Role == "end" {
			end = Rooms[i]
			break
		}
	}

	if end == nil || len(end.Parent) == 0 {
		return nil
	}

	// Récupérer tous les chemins avec backtracking récursif
	var allPaths [][]*Room
	var currentPath []*Room
	findAllPathsRecursive(end, currentPath, &allPaths)

	return allPaths
}

func findAllPathsRecursive(room *Room, currentPath []*Room, allPaths *[][]*Room) {
	// Ajouter la salle actuelle au chemin
	currentPath = append([]*Room{room}, currentPath...)

	// Si on arrive au start, on a trouvé un chemin complet
	if room.Role == "start" {
		// Copier le chemin pour éviter les modifications
		pathCopy := make([]*Room, len(currentPath))
		copy(pathCopy, currentPath)
		*allPaths = append(*allPaths, pathCopy)
		return
	}

	// Explorer tous les parents
	for _, parent := range room.Parent {
		findAllPathsRecursive(parent, currentPath, allPaths)
	}
}
