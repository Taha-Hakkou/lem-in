package lemin

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(lines []string) (int, []*Room, []Link, error) {
	if len(lines) == 0 {
		return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
	}

	indexAnts := -1
	for i, r := range lines {
		if strings.HasPrefix(r, "#") && r != "##start" && r != "##end" {
			continue
		}
		indexAnts = i
		break
	}
	if indexAnts == -1 {
		indexAnts = 0
	}

	Ants, err := strconv.Atoi(lines[indexAnts])
	if err != nil || Ants < 1 {
		return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, invalid number of Ants")
	}

	var rooms []string
	var links []string
	var startCount int
	var endCount int
	lines = lines[indexAnts+1:] // FIX: Skip the ants line

	for i, line := range lines {
		line = strings.TrimSpace(line)
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

	if startCount != 1 {
		return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, no start room found")
	}
	if endCount != 1 {
		return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, no end room found")
	}

	var Rooms []Room
	var CheckingCordin [][]string
	for i := 0; i < len(rooms); i++ {
		r := rooms[i]
		fields := strings.Fields(r)

		// Skip comments (but not ##start and ##end)
		if len(fields) == 1 && strings.HasPrefix(r, "#") && r != "##start" && r != "##end" {
			continue
		}

		// FIX: Add parentheses for correct operator precedence
		if len(fields) == 3 && (strings.HasPrefix(fields[0], "#") || strings.HasPrefix(fields[0], "L") || strings.Contains(fields[0], " ")) {
			return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, invalid room name")
		}

		if r == "##start" || r == "##end" {
			if i+1 >= len(rooms) {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
			}
			next := strings.Fields(rooms[i+1])
			if len(next) != 3 {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
			}

			// FIX: Check room name validity for start/end rooms
			if strings.HasPrefix(next[0], "#") || strings.HasPrefix(next[0], "L") || strings.Contains(next[0], " ") {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, invalid room name")
			}

			x, err := strconv.Atoi(next[1])
			if err != nil {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, invalid coordinates")
			}
			y, err := strconv.Atoi(next[2])
			if err != nil {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, invalid coordinates")
			}
			if x < 0 || y < 0 {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, negative coordinates")
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
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, duplicate room")
			}
			Rooms = append(Rooms, newRoom)
			i++ // skip the room line after ##start/##end
			continue
		}

		// Handle normal rooms (NOT preceded by ##start or ##end)
		if len(fields) == 3 {
			tempLink := append([]string{}, fields[1], fields[2])
			if checkRepCoor(CheckingCordin, tempLink) {
				CheckingCordin = append(CheckingCordin, tempLink)
			} else {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, duplicate coordinate")
			}
			x, err := strconv.Atoi(fields[1])
			if err != nil {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, invalid coordinates")
			}
			y, err := strconv.Atoi(fields[2])
			if err != nil {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, invalid coordinates")
			}
			if x < 0 || y < 0 {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, negative coordinates")
			}
			newRoom := Room{Name: fields[0], X: x, Y: y, Role: "normal"}
			if !checkRepRooms(Rooms, fields[0]) {
				return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, duplicate room")
			}
			Rooms = append(Rooms, newRoom)
		}
	}

	var Checking [][]string
	for _, link := range links {
		tempLink := strings.Split(link, "-")
		if strings.HasPrefix(link, "#") {
			continue
		}
		if len(tempLink) != 2 || tempLink[0] == "" || tempLink[1] == "" {
			return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, invalid link")
		}
		if checkRepRooms(Rooms, tempLink[0]) || checkRepRooms(Rooms, tempLink[1]) {
			return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, link to unknown room")
		}
		if tempLink[0] == tempLink[1] {
			return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, room links to itself")
		}
		if checkRepLinks(Checking, tempLink) {
			Checking = append(Checking, tempLink)
		} else {
			return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, duplicate link")
		}
	}

	var parsedRooms []*Room
	for i := range Rooms {
		parsedRooms = append(parsedRooms, &Rooms[i])
	}

	var Links []Link
	for _, l := range Checking {
		r1 := FindRoom(parsedRooms, l[0])
		r2 := FindRoom(parsedRooms, l[1])
		if r1 == nil || r2 == nil {
			return 0, nil, nil, fmt.Errorf("ERROR: invalid data format, link to unknown room")
		}
		Links = append(Links, Link{R1: r1, R2: r2})
	}

	return Ants, parsedRooms, Links, nil
}
