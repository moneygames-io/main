package main

type Tile struct {
	Snake *Snake
	Depth int
}

type Map struct {
	Tiles [][]Tile
}

func (m *Map) NewMap(width int, height int) *Map {
	newMap := new(Map)
	//TODO
	return newMap
}

func (m *Map) spawnNewSnake(Player string, initialLen int, startingX int, startingY int) {

}
