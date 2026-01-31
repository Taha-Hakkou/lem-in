package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(lines []string) (int, []Room, []Link, error) {
	if len(lines) == 0 {
		return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
	}
	Ants, err := strconv.Atoi(lines[0])
	if err != nil || Ants < 1 {
		return 0, nil, nil, fmt.Errorf("ERROR: invalid number of Ants")
	}
	var rooms []string
	var links []string
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
			return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
		}
		if len(Templine) == 1 && !strings.Contains(line, "-") && !strings.HasPrefix(line, "#") {
			return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")

		}
		if strings.Contains(line, "-") {
			index := i
			rooms = lines[:index]
			links = lines[index:]
			break
		}
	}
	if startCount != 1 || endCount != 1 {
		return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
	}

	var Rooms []Room
	for i := 0; i < len(rooms); i++ {
		r := rooms[i]
		fields := strings.Fields(r)

		// Skip comments (but not ##start and ##end)
		if len(fields) == 1 && strings.HasPrefix(r, "#") && r != "##start" && r != "##end" {
			continue
		}
		if len(fields) == 3 && strings.HasPrefix(fields[0], "#") || strings.HasPrefix(fields[0], "L") {
			return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
		}

		if r == "##start" || r == "##end" {
			if i+1 >= len(rooms) {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
			}
			next := strings.Fields(rooms[i+1])
			if len(next) != 3 {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
			}
			x, err := strconv.Atoi(next[1])
			if err != nil {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
			}
			y, err := strconv.Atoi(next[2])
			if err != nil {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
			}
			role := ""
			if r == "##start" {
				role = "start"
			}
			if r == "##end" {
				role = "end"
			}

			newRoom := Room{Name: next[0], X: x, Y: y, Role: role}
			if !checkRepRooms(Rooms, next[0]) {
				return 0, nil, nil, fmt.Errorf("ERROR: duplicate room")
			}
			Rooms = append(Rooms, newRoom)
			i++ // skip the room line after ##start/##end
			continue
		}

		// Handle normal rooms (NOT preceded by ##start or ##end)
		if len(fields) == 3 {
			x, err := strconv.Atoi(fields[1])
			if err != nil {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
			}
			y, err := strconv.Atoi(fields[2])
			if err != nil {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
			}
			newRoom := Room{Name: fields[0], X: x, Y: y, Role: "normal"}
			if !checkRepRooms(Rooms, fields[0]) {
				return 0, nil, nil, fmt.Errorf("ERROR: duplicate room")
			}
			Rooms = append(Rooms, newRoom)
		}
	}
	var Checking []string
	for _, link := range links {
		tempLink := strings.Split(link, "-")
		if strings.HasPrefix(link, "#") && link != "##start" && link != "##end" {
			continue
		}
		if len(tempLink) != 2 {
			return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
		}
		if checkRepRooms(Rooms, tempLink[0]) || checkRepRooms(Rooms, tempLink[1]) || tempLink[0] == tempLink[1] {
			return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
		}
		if checkRepLinks(Checking, link) {
			Checking = append(Checking, link)
		} else {
			return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
		}

	}
	var Links []Link
	for _, l := range Checking {
		splited := strings.Split(l, "-")
		Links = append(Links, Link{R1: FindRoom(Rooms, splited[0]), R2: FindRoom(Rooms, splited[1])})
	}

	return Ants, Rooms, Links, nil
}
