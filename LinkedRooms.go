package main

func GetRelatedRooms(Rooms []Room, Links []Link) []Room {
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
	return Rooms
}
