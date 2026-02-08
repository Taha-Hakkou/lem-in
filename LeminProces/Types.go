package lemin

type Room struct {
	Name      string
	X         int
	Y         int
	Role      string // "start", "end", "normal"
	Relations []*Room
}

type Ant struct {
	Pos    int
	Number int
	Path   []*Room
}

type Link struct {
	R1 *Room
	R2 *Room
}
