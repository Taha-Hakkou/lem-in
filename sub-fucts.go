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
func FindRoom(Rooms []Room, name string) *Room {
	for _, r := range Rooms {
		if r.Name == name {
			return &r
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
	Rooms = GetRelatedRooms(Rooms, Links)
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
		fmt.Printf("Room %s -> %s\n", r.Name, r.PrintRelations())
	}
}
