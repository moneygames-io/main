package main

import (
	"math/rand"
	"github.com/Parth/boolean"
)

type Tile struct {
	Snake *Snake
	Food *Food
}

type Map struct {
	Tiles [][]Tile
	Players map[*Player]*Snake
}

type MapEvent interface {
	SnakeCreated(*Snake)
	AddNode(*SnakeNode) int
	RemoveNode(int, int)
	AddFood(*Food)
	RemoveFood(int, int)
	SnakeRemoved(*Snake)
}

func NewMap(players int) *Map {
	newMap := &Map{}
	newMap.Tiles = make([][]Tile, players * 10)

	for i := range newMap.Tiles {
		newMap.Tiles[i] = make([]Tile, players * 10)
	}

	newMap.Players = make(map[*Player]*Snake)
	return newMap
}

func (m *Map) SpawnNewPlayer(player *Player) {
	row := rand.Intn(len(m.Tiles))
	col := rand.Intn(len(m.Tiles[0]))

	for m.Tiles[row][col].Snake != nil && m.Tiles[row][col].Food != nil {
		row = rand.Intn(len(m.Tiles))
		col = rand.Intn(len(m.Tiles[0]))
	}

	m.Players[player] = NewSnake(col, row, m)
}

func (m *Map) SnakeCreated(snake *Snake) {
	m.AddNode(snake.Head)
}

func (m *Map) AddNode(snakeNode *SnakeNode) int {
	col := snakeNode.X
	row := snakeNode.Y

	if m.Tiles[row][col].Snake != nil {
		if snakeNode.Snake != m.Tiles[row][col].Snake {
			return 2
		}
	}

	m.Tiles[row][col].Snake = snakeNode.Snake
	return boolean.BtoI(m.Tiles[row][col].Food != nil)
}

func (m *Map) RemoveNode(col int, row int) {
	m.Tiles[row][col].Snake = nil
}

func (m *Map) AddFood(food *Food) {
	col := food.X
    row := food.Y

	m.Tiles[row][col].Food = food
}

func (m *Map) RemoveFood(col int, row int) {
	m.Tiles[row][col].Food = nil

}

func (m *Map) SnakeRemoved(snake *Snake) {
	m.Players[snake.Player] = nil
}

func (m *Map) update() {
	for player, snake := range m.Players {
		if player.CurrentSprint {
			snake.Sprint(player.CurrentDirection)
		} else {
			snake.Move(player.CurrentDirection)
		}
	}
}

func (m *Map) render() [][]Tile {
	return m.Tiles
}
