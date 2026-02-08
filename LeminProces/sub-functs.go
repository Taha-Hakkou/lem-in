package lemin

func checkRepRooms(sl []Room, s string) bool {
	for _, r := range sl {
		if r.Name == s {
			return false
		}
	}
	return true
}

func checkRepLinks(sl []string, s string) bool {
	for _, r := range sl {
		if r == s {
			return false
		}
	}
	return true
}

func FindRoom(rooms []*Room, name string) *Room {
	for i := range rooms {
		if rooms[i].Name == name {
			return rooms[i]
		}
	}
	return nil
}

func Clear(Lines []string) []string {
	var result []string
	for _, r := range Lines {
		if r != "" {
			result = append(result, r)
		}
	}
	return result
}

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
