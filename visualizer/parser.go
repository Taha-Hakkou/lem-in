package main

import (
	"regexp"
	"strconv"
	"strings"
)

var start, end string

func parseData(lines []string) {
	// Number of ants
	// n, _ = strconv.Atoi(lines[0])

	var i int

	// Room parsing
	roomex := regexp.MustCompile(`^([A-Za-z0-9]+) ([0-9]+) ([0-9]+)$`)
	for i = 1; i < len(lines); i++ {
		if roomex.MatchString(lines[i]) { // contains and not match: resolved
			groups := roomex.FindStringSubmatch(lines[i])
			x, _ := strconv.Atoi(groups[2])
			y, _ := strconv.Atoi(groups[3])
			rooms[groups[1]] = Room{x, y}
			// fmt.Println(lines[i])
		} else if lines[i] == "##start" {
			start = strings.Split(lines[i+1], " ")[0]
		} else if lines[i] == "##end" {
			end = strings.Split(lines[i+1], " ")[0]
		} else if strings.HasPrefix(lines[i], "#") {
			continue
		} else {
			break
		}
	}

	// Tunel parsing
	tunlex := regexp.MustCompile(`^([A-Za-z0-9]+)-([A-Za-z0-9]+)$`)
	for ; i < len(lines); i++ {
		if tunlex.MatchString(lines[i]) { // contains and not match: resolved
			groups := tunlex.FindStringSubmatch(lines[i])
			links = append(links, [2]string{groups[1], groups[2]})
			// fmt.Println(lines[i])
		} else {
			break
		}
	}

	// Link parsing
	i++ // skip empty line
	stepx := regexp.MustCompile(`L([A-Za-z0-9]+)-([A-Za-z0-9]+)`)
	for ; i < len(lines); i++ {
		// no need to check if matching anymore
		moves := stepx.FindAllStringSubmatch(lines[i], -1)
		newMoves := make([][2]string, len(moves))
		for j, move := range moves {
			newMoves[j] = [2]string{move[1], move[2]}
		}
		steps = append(steps, newMoves)
	}
}
