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

	ants, err := strconv.Atoi(lines[0])
	if err != nil || ants < 1 {
		return 0, nil, nil, fmt.Errorf("ERROR: invalid number of ants")
	}

	lines = Clear(lines[1:])

	var roomsDef []string
	var linksDef []string
	startCount, endCount := 0, 0

	for i, line := range lines {
		if line == "##start" {
			startCount++
		}
		if line == "##end" {
			endCount++
		}

		if strings.Contains(line, "-") && !strings.HasPrefix(line, "#") {
			roomsDef = lines[:i]
			linksDef = lines[i:]
			break
		}
	}

	if startCount != 1 || endCount != 1 {
		return 0, nil, nil, fmt.Errorf("ERROR: invalid data format")
	}

	var rooms []Room

	for i := 0; i < len(roomsDef); i++ {
		line := roomsDef[i]
		if strings.HasPrefix(line, "#") && line != "##start" && line != "##end" {
			continue
		}

		if line == "##start" || line == "##end" {
			role := "start"
			if line == "##end" {
				role = "end"
			}
			i++
			fields := strings.Fields(roomsDef[i])
			x, _ := strconv.Atoi(fields[1])
			y, _ := strconv.Atoi(fields[2])

			if !checkRepRooms(rooms, fields[0]) {
				return 0, nil, nil, fmt.Errorf("ERROR: duplicate room")
			}

			rooms = append(rooms, Room{
				Name: fields[0],
				X:    x,
				Y:    y,
				Role: role,
			})
			continue
		}

		fields := strings.Fields(line)
		x, _ := strconv.Atoi(fields[1])
		y, _ := strconv.Atoi(fields[2])

		if !checkRepRooms(rooms, fields[0]) {
			return 0, nil, nil, fmt.Errorf("ERROR: duplicate room")
		}

		rooms = append(rooms, Room{
			Name: fields[0],
			X:    x,
			Y:    y,
			Role: "normal",
		})
	}

	var roomPtrs []*Room
	for i := range rooms {
		roomPtrs = append(roomPtrs, &rooms[i])
	}

	var links []Link
	var seen []string

	for _, l := range linksDef {
		if strings.HasPrefix(l, "#") {
			continue
		}
		if !checkRepLinks(seen, l) {
			return 0, nil, nil, fmt.Errorf("ERROR: duplicate link")
		}
		seen = append(seen, l)

		s := strings.Split(l, "-")
		r1 := FindRoom(roomPtrs, s[0])
		r2 := FindRoom(roomPtrs, s[1])
		if r1 == nil || r2 == nil {
			return 0, nil, nil, fmt.Errorf("ERROR: invalid link")
		}
		links = append(links, Link{r1, r2})
	}

	return ants, roomPtrs, links, nil
}
