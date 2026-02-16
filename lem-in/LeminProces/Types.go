package lemin

type Room struct {
	Name      string
	X, Y      int
	Role      string // "start", "end", "normal"
	Relations []*Room
}

type Link struct {
	R1 *Room
	R2 *Room
}

type Ant struct {
	Number int
	Pos    int
	Path   []*Room
}
