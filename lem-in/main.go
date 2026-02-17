package main

import (
	"fmt"
	"os"
	"strings"

	lemin "main/LeminProces"
)

func main() {
	// Validate command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./lem-in [FILE]")
		return
	}

	// Read input file
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Normalize line endings and split into lines
	bytes = []byte(strings.ReplaceAll(string(bytes), "\r\n", "\n"))
	lines := strings.Split(string(bytes), "\n")

	// Parse input: extract ants, rooms, and links
	Ants, Rooms, Links, err := lemin.Parse(lemin.Clear(lines))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Build room relationships (adjacency list)
	lemin.GetRelatedRooms(Rooms, Links)

	// Find optimal paths from start to end
	paths := lemin.PathFinder(Rooms, Links)

	if len(paths) == 0 {
		fmt.Println("ERROR: no path found")
		return
	}

	// Print original input
	fmt.Println(strings.Join(lemin.Clear(lines), "\n"))
	fmt.Println()

	// Simulate ant movement and print results
	lemin.MoveAnts(paths, Ants)
}