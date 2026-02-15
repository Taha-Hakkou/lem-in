package main

import (
	"fmt"
	"os"
	"strings"

	lemin "main/LeminProces"
)

func Print(Args []string) {
	if len(Args) != 2 || len(Args[1]) <= 4 || !strings.HasSuffix(Args[1], ".txt") {
		fmt.Println("Usage: ./lem-in [FILE]")
		return
	}
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes = []byte(strings.ReplaceAll(string(bytes), "\r\n", "\n"))
	lines := strings.Split(string(bytes), "\n")
	Ants, Rooms, Links, err := lemin.Parse(lemin.Clear(lines))
	if err != nil {
		fmt.Println(err)
		return
	}

	lemin.GetRelatedRooms(Rooms, Links)

	// Use intelligent path finding
	paths := lemin.PathFinder(Rooms, Links)

	if len(paths) == 0 {
		fmt.Println("ERROR: no path found")
		return
	}
	// Move ants
	lemin.MoveAnts(paths, Ants)
}
