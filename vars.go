package main

import (
	"fmt"
)

type Room struct {
	Name      string
	X         int
	Y         int
	Role      string // "start", "end", "normal"
	Relations []*Room


}

func (r *Room) PrintRooms() string {
	return fmt.Sprintf("%s (%d,%d) [%s]", r.Name, r.X, r.Y, r.Role)
}
func (r *Room) PrintRelations() string {
	var out string
	for _, rel := range r.Relations {
		out += rel.Name + " "
	}
	return out
}

type Link struct {
	R1 *Room
	R2 *Room
}

func (l *Link) PrintLinks() string {
	return fmt.Sprintf("%s [%s] - %s [%s]", l.R1.Name, l.R1.Role, l.R2.Name, l.R2.Role)
}
