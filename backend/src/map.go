// TODO Standardize between row/col and x/y across the app
package main

import (
	"github.com/Parth/boolean"
	"math/rand"
	"time"
)

type Tile struct {
	Snake *Snake
	Food  *Food
}

type Map struct {
	Tiles   [][]Tile
	Players map[*Player]*Snake
	Losers  map[*Player]*Snake
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
	newMap.Tiles = make([][]Tile, players*10)

	for i := range newMap.Tiles {
		newMap.Tiles[i] = make([]Tile, players*10)
	}

	newMap.Players = make(map[*Player]*Snake)
	newMap.Losers = make(map[*Player]*Snake)
	return newMap
}

func (m *Map) SpawnNewPlayer(player *Player) (int, int) {
	rand.Seed(time.Now().UnixNano())
	row := rand.Intn(len(m.Tiles))
	col := rand.Intn(len(m.Tiles[0]))

	for m.Tiles[row][col].Snake != nil && m.Tiles[row][col].Food != nil {
		// TODO infinite loop risk
		row = rand.Intn(len(m.Tiles))
		col = rand.Intn(len(m.Tiles[0]))
	}

	return m.SpawnNewPlayerAt(player, col, row)
}

func (m *Map) SpawnNewPlayerAt(player *Player, col int, row int) (int, int) {
	// TODO check if occupied?
	m.Players[player] = NewSnake(col, row, m, player)
	player.Snake = m.Players[player]

	// TODO is this return needed? 
	// Possibly, could be used to indicate a different location if this one is occupied
	return col, row
}

func (m *Map) SnakeCreated(snake *Snake) {
	m.AddNode(snake.Head)
}

func (m *Map) AddNode(snakeNode *SnakeNode) int {
	col := snakeNode.X
	row := snakeNode.Y

	if row >= len(m.Tiles) || col >= len(m.Tiles[0]) {
		return 2
	}

	if row < 0 || col < 0 {
		return 2
	}

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
	m.Losers[snake.Player] = snake
}

func (m *Map) Update() {
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
