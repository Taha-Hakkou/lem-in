package main

import (
	"fmt"
	"os"
	"strings"

	lemin "main/LeminProces"
)

func main() {
	if len(os.Args) != 2 || len(os.Args[1]) <= 4 || !strings.HasSuffix(os.Args[1], ".txt") {
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

	paths := lemin.PathFinder(Rooms, Links)

	if len(paths) == 0 {
		fmt.Println("ERROR: no path found")
		return
	}
	fmt.Println(strings.Join(lemin.Clear(lines), "\n"))
	fmt.Println()

	lemin.MoveAnts(paths, Ants)
}
