type Tile struct {
	Snake Snake*
	Depth int
}

type Map struct {
	Tiles [][]Tile
}
