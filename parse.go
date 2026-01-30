package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(lines []string) (int, error) {
	Ants, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, fmt.Errorf("ERROR: invalid number of Ants")
	}
	var rooms []string
	// TODO: var links []string
	var startCount int
	var endCount int
	lines = lines[1:]
	for i, line := range lines {
		Templine := strings.Fields(line)
		if line == "##start" {
			startCount++
		}
		if line == "##end" {
			endCount++
		}
		if len(Templine) > 3 || len(Templine) < 1 {
			return 0, fmt.Errorf("ERROR: invalid data format")
		}
		if len(Templine) == 1 && !strings.Contains(line, "-") && !strings.HasPrefix(line, "#") {
			return 0, fmt.Errorf("ERROR: invalid data format")

		}
		if strings.Contains(line, "-") {
			index := i
			rooms = lines[:index]
			// TODO: links = lines[index:]
			break
		}
	}
	if startCount != 1 || endCount != 1 {
		return 0, fmt.Errorf("ERROR: invalid data format")
	}

	var Rooms []Room
	for i := 0; i < len(rooms); i++ {
		r := rooms[i]

		// Skip comments (but not ##start and ##end)
		if strings.HasPrefix(r, "#") && r != "##start" && r != "##end" {
			continue
		}

		if r == "##start" || r == "##end" {
			if i+1 >= len(rooms) {
				return 0, fmt.Errorf("ERROR: invalid data format")
			}
			next := strings.Fields(rooms[i+1])
			if len(next) != 3 {
				return 0, fmt.Errorf("ERROR: invalid data format")
			}
			x, err := strconv.Atoi(next[1])
			if err != nil {
				return 0, fmt.Errorf("ERROR: invalid data format")
			}
			y, err := strconv.Atoi(next[2])
			if err != nil {
				return 0, fmt.Errorf("ERROR: invalid data format")
			}
			role := ""
			if r == "##start" {
				role = "start"
			}
			if r == "##end" {
				role = "end"
			}
			newRoom := Room{name: next[0], x: x, y: y, role: role}
			if !checkRep(Rooms, newRoom) {
				return 0, fmt.Errorf("ERROR: duplicate room")
			}
			Rooms = append(Rooms, newRoom)
			i++ // skip the room line after ##start/##end
			continue
		}

		// Handle normal rooms (NOT preceded by ##start or ##end)
		fields := strings.Fields(r)
		if len(fields) == 3 {
			x, err := strconv.Atoi(fields[1])
			if err != nil {
				return 0, fmt.Errorf("ERROR: invalid data format")
			}
			y, err := strconv.Atoi(fields[2])
			if err != nil {
				return 0, fmt.Errorf("ERROR: invalid data format")
			}
			newRoom := Room{name: fields[0], x: x, y: y, role: "normal"}
			if !checkRep(Rooms, newRoom) {
				return 0, fmt.Errorf("ERROR: duplicate room")
			}
			Rooms = append(Rooms, newRoom)
		}
	}

	fmt.Println(Rooms)

	return Ants, nil
}
