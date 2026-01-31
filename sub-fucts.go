package main

import "fmt"

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
	pathFound := FindPath(Rooms)

	if !pathFound {
		fmt.Println("ERROR: no path found")
		return
	}

	fmt.Println("Number of Ants :", Ants)
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
		parentName := "nil"
		if r.Parent != nil {
			parentName = r.Parent.Name
		}
		fmt.Printf("Room %s -> parent name: %s -> steps: %d\n", r.Name, parentName, r.Steps)
	}
	fmt.Println()

	Path := ExtractPaths(Rooms)
	if Path == nil {
		fmt.Println("ERROR: no path found")
		return
	}

	for i := len(Path) - 1; i >= 0; i-- {
		fmt.Print(Path[i].Name)
		if i != 0 {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
