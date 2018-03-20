package main

import (
	"math/rand"
)

type Tile struct {
	Snake *Snake
	Food *Food
}

type Map struct {
	Tiles [][]Tile
	Players map[string]*Snake
	Losers map[string]*Snake
}

type MapEvent interface {
	SnakeCreated(*Snake)
	HeadMoved(*Snake) int
	RemoveTailNode(*SnakeNode)
	AddFood(*Food)
	RemoveFood(*Food)
	SnakeRemoved(*Snake)
}

func NewMap(players int) *Map {
	newMap := &Map{}
	newMap.Tiles = [players * 100][players * 100]Tile{}
	newMap.Players = make(map[string]*Snake)
	newMap.Losers = make(map[string]*Snake)
	return newMap
}

func (m *Map) SpawnNewPlayer(player string) bool {
	row := rand.Intn(len(m.Tiles))
	col := rand.Intn(len(m.Tiles[0]))

	for m.Tiles[row][col].Snake != nil && m.Tiles[row][col].Food != nil {
		row = rand.Intn(len(m.Tiles))
		col = rand.Intn(len(m.Tiles[0]))
	}

	map.Players[player] = NewSnake(col, row, m)
}

func (m *Map) SnakeCreated(snake *Snake) {
	m.Tiles[x][y].Snake = snake
	m.Tiles[x][y].Depth++
}

func (m *Map) HeadMoved(snake *Snake) int {
	x := snake.Head.X
	y := snake.Head.Y

	if m.Tiles[y][x].Snake != nil {
		if snake != m.Tiles[y][x] {
			return 2
		}
	}

	if m.Tiles[y][x].Food != nil {
		m.RemoveFood(m.Tiles[y][x].Food)
		return 1
	}

	return 0
}

func (m *Map) RemoveTailNode(sn *SnakeNode) {
	x := sn.X
    y := sn.Y

	m.Tiles[y][x].Snake = nil
}

func (m *Map) AddFood(food *Food) {
	x := food.X
    y := food.Y

	m.Tiles[y][x] = food
}

func (m *Map) RemoveFood(food *Food) {
	x := food.X
    y := food.Y

	m.Tiles[y][x].Food = nil

}

func (m *Map) SnakeRemoved(snake *Snake) {
	m.Players[snake.Player] = nil
	m.Losers[snake.Player] = snake
}

func (m *Map) tick() {

}
