package main

import (
	"fmt"
)

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

func PrintParsedData(lines []string) {
	Ants, Rooms, Links, err := Parse(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	GetRelatedRooms(Rooms, Links)

	fmt.Println("Number of Ants:", Ants)
	fmt.Println()

	// Extraire et afficher tous les chemins
	allPaths := FindAllPaths(Rooms)
	if allPaths == nil {
		fmt.Println("ERROR: no path found")
		return
	}
	

	MoveAnts(allPaths, Ants)
}
