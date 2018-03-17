package main

type Tile struct {
	Snake *Snake
	Depth int
}

type Map struct {
	Tiles [][]Tile
}

type MapEvent interface {
	AddSnakeNode(*SnakeNode)
	RemoveSnakeNode(*SnakeNode)
	AddFood(*Food)
	RemoveFood(*Food)
}

func (m *Map) NewMap(width int, height int) *Map {
	newMap := new(Map)
	//TODO
	return newMap
}

func (m *Map) AddSnakeNode(snake *SnakeNode) {

}

func (m *Map) RemoveSnakeNode(snake *SnakeNode) {

}

func (m *Map) AddFood(*Food) {

}

func (m *Map) RemoveFood(*Food) {

}
