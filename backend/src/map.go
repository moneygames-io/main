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
	Players map[string]*Player
	Losers map[string]*Player
}

type MapEvent interface {
	SnakeCreated(*Snake)
	AddNode(*SnakeNode) int
	RemoveNode(row int, col int)
	AddFood(*Food)
	RemoveFood(row int, col int)
	SnakeRemoved(*Snake)
}

func NewMap(players int) *Map {
	newMap := &Map{}
	newMap.Tiles = [players * 100][players * 100]Tile{}
	newMap.Players = make(map[string]*Player)
	newMap.Losers = make(map[string]*Player)
	return newMap
}

func (m *Map) SpawnNewPlayer(player *Player) bool {
	row := rand.Intn(len(m.Tiles))
	col := rand.Intn(len(m.Tiles[0]))

	for m.Tiles[row][col].Snake != nil && m.Tiles[row][col].Food != nil {
		row = rand.Intn(len(m.Tiles))
		col = rand.Intn(len(m.Tiles[0]))
	}

	map.Players[player.Name] = NewSnake(col, row, m)
}

func (m *Map) SnakeCreated(snake *Snake) {
	m.Tiles[x][y].Snake = snake
	m.Tiles[x][y].Depth++
}

func (m *Map) AddNode(snakeNode *SnakeNode) int {
	x := snakeNode.X
	y := snakeNode.Y

	if m.Tiles[y][x].Snake != nil {
		if snake != m.Tiles[y][x] {
			return 2
		}
	}

	if m.Tiles[y][x].Food != nil {
		return 1
	}

	return 0
}

func (m *Map) RemoveTailNode(x int, y int) {
	m.Tiles[y][x].Snake = nil
}

func (m *Map) AddFood(food *Food) {
	x := food.X
    y := food.Y

	m.Tiles[y][x] = food
}

func (m *Map) RemoveFood(x int, y int) {
	m.Tiles[y][x].Food = nil

}

func (m *Map) SnakeRemoved(snake *Snake) {
	m.Players[snake.Player] = nil
	m.Losers[snake.Player] = snake
}

func (m *Map) update() {

}

func (m *Map) render() [][]Tile {

}
