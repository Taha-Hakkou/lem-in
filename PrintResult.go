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
	Ants, Rooms, Links, err := lemin.Parse(lines)
	if err != nil {
		fmt.Println(err)
		return
	}
	lemin.GetRelatedRooms(Rooms, Links)

	allPaths := lemin.FindAllPaths(Rooms)
	if allPaths == nil {
		fmt.Println("ERROR: no path found")
		return
	}

	lemin.MoveAnts(allPaths, Ants)
}

