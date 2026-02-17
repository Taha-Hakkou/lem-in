package lemin

// checkRepRooms returns true if room name doesn't exist (no duplicate)
func checkRepRooms(sl []Room, s string) bool {
	for _, r := range sl {
		if r.Name == s {
			return false
		}
	}
	return true
}

// checkRepLinks returns true if link doesn't exist (checks both directions)
func checkRepLinks(sl [][]string, s []string) bool {
	for _, r := range sl {
		// Check both "A-B" and "B-A" as duplicates
		if (r[0] == s[0] && r[1] == s[1]) || (r[0] == s[1] && r[1] == s[0]) {
			return false
		}
	}
	return true
}

// checkRepCoor returns true if coordinates don't exist (no duplicate)
func checkRepCoor(sl [][]string, s []string) bool {
	for _, r := range sl {
		if r[0] == s[0] && r[1] == s[1] {
			return false
		}
	}
	return true
}

// FindRoom returns pointer to room with given name, or nil if not found
func FindRoom(rooms []*Room, name string) *Room {
	for i := range rooms {
		if rooms[i].Name == name {
			return rooms[i]
		}
	}
	return nil
}

// Clear removes empty lines from input
func Clear(Lines []string) []string {
	var result []string
	for _, r := range Lines {
		if r != "" {
			result = append(result, r)
		}
	}
	return result
}

// GetRelatedRooms populates each room's Relations list with its neighbors
func GetRelatedRooms(Rooms []*Room, Links []Link) {
	for i := range Rooms {
		for _, l := range Links {
			// Add relationships
			if Rooms[i].Name == l.R1.Name {
				Rooms[i].Relations = append(Rooms[i].Relations, l.R2)
			}
			if Rooms[i].Name == l.R2.Name {
				Rooms[i].Relations = append(Rooms[i].Relations, l.R1)
			}
		}
	}
}