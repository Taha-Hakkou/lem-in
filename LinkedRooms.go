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

	for len(Queue) > 0 {
		Curent := Queue[0]
		Queue = Queue[1:]
		for i := range Curent.Relations {
			curentTemp := Curent.Relations[i]
			if curentTemp.IsVisited {
				continue
			}
			curentTemp.IsVisited = true
			curentTemp.Parent = Curent
			curentTemp.Steps = Curent.Steps + 1
			if curentTemp.Role == "end" {
				return true // Path found
			}
			Queue = append(Queue, curentTemp)
		}
	}
	return false 
}

func ExtractPaths(Rooms []*Room) []*Room {
	var End *Room
	var Path []*Room
	for i := range Rooms {
		if Rooms[i].Role == "end" {
			End = Rooms[i]
			break
		}
	}

	if End == nil || End.Parent == nil {
		return nil
	}

	Path = append(Path, End)
	current := End

	for current.Parent != nil {
		Path = append(Path, current.Parent)
		current = current.Parent
	}

	return Path
}
