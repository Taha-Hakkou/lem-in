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

	for _, r := range Rooms {
		fmt.Println(r.PrintRooms())
	}
	fmt.Println()

	for _, l := range Links {
		fmt.Println(l.PrintLinks())
	}
	fmt.Println()
	for _, r := range Rooms {
		fmt.Println(r.Name + "->" + r.PrintRelations())
	}
	fmt.Println()

	// Extraire et afficher tous les chemins
	allPaths := FindAllPaths(Rooms)
	if allPaths == nil {
		fmt.Println("ERROR: no path found")
		return
	}

	fmt.Printf("Found %d path(s):\n", len(allPaths))
	for i, path := range allPaths {
		fmt.Printf("Path %d: ", i+1)
		for j, room := range path {
			fmt.Print(room.Name)
			if j < len(path)-1 {
				fmt.Print(" -> ")
			}
		}
		fmt.Printf(" (length: %d)\n", len(path)-1)
	}
	MoveAnts(allPaths ,Ants)
}
