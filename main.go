package main

import (
	"fmt"
	"os"
	"strings"
)

// quickest path: means quickest by rounds, number of tunels traversedm OR time ?!
// => as few moves as possible.
// The shortest path is not necessarily the simplest.

// Some colonies will have many rooms and many links, but no path between ##start and ##end.
// Some will have rooms that link to themselves, sending your path-search spinning in circles. Some will have too many/too few ants, no ##start or ##end, duplicated rooms, links to unknown rooms, rooms with invalid coordinates and a variety of other invalid or poorly-formatted input. In those cases the program will return an error message ERROR: invalid data format. If you wish, you can elaborate a more specific error message (example: ERROR: invalid data format, invalid number of Ants or ERROR: invalid data format, no start room found).

// A room will never start with the letter L or with # and must have no spaces.
// You join the rooms together with as many tunnels as you need.
// A tunnel joins only two rooms together never more than that.
// A room can be linked to multiple rooms.
// Two rooms can't have more than one tunnel connecting them.
// Each room can only contain one ant at a time (except at ##start and ##end which can contain as many ants as necessary).
// Each tunnel can only be used once per turn.
// To be the first to arrive, ants will need to take the shortest path or paths. They will also need to avoid traffic jams as well as walking all over their fellow ants.
// You will only display the ants that moved at each turn, and you can move each ant only once and through a tunnel (the room at the receiving end must be empty).
// The rooms names will not necessarily be numbers, and in order.
// Any unknown command will be ignored.
// The program must handle errors carefully. In no way can it quit in an unexpected manner.
// The coordinates of the rooms will always be int.

func main() {
	if len(os.Args) != 2 || len(os.Args[1]) <= 4 || !strings.HasSuffix(os.Args[1], ".txt") {
		fmt.Println("Usage: ./lem-in [FILE]")
		return
	}
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	bytes = []byte(strings.ReplaceAll(string(bytes), "\r\n", "\n"))
	lines := strings.Split(string(bytes), "\n")
	fmt.Println(Parse(lines))

}
