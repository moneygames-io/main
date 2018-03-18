package main

type Tile struct {
	Snake *Snake
	Depth int
}

type Map struct {
	Tiles [][]Tile
}

type MapEvent interface {
	SnakeCreated(*Snake)
	HeadMoved(*Snake) int
	RemoveTailNode(*SnakeNode)
	AddFood(*Food)
	RemoveFood(*Food)
	SnakeRemoved(*Snake)
}

func (m *Map) NewMap(width int, height int) *Map {
	newMap := new(Map)
	// TODO
	return newMap
}
