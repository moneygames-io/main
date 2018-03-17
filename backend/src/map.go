package main

type Tile struct {
	Snake *Snake
	Depth int
}

type Map struct {
	Tiles [][]Tile
}

type MapEvent interface {
	AddSnakeNode(snakeNode SnakeNode*)
	RemoveSnakeNode(snakeNode SnakeNode*)
	AddFood(food Food*)
	RemoveFood(food Food*)
}

func (m *Map) NewMap(width int, height int) *Map {
	newMap := new(Map)
	//TODO
	return newMap
}

func (m *Map) AddSnakeNode(snake SnakeNode*) {

}

func (m *Map) RemoveSnakeNode(snake SnakeNode*) {

}

func (m *Map) AddFood(x int, y int) {

}

func (m *Map) RemoveFood(player string) {

}
