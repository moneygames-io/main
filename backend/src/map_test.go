package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func printMap(m *Map) {
	for r := range m.Tiles {
		for c := range m.Tiles[0] {
			if m.Tiles[r][c].Snake != nil {
				fmt.Print("S")
			} else if m.Tiles[r][c].Food != nil {
				fmt.Print("F")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func TestNewMap(t *testing.T) {
	m := NewMap(1)
	assert.NotNil(t, m)
	assert.NotNil(t, m.Tiles)
	assert.NotNil(t, m.Players)
	for r := 0; r < len(m.Tiles); r++ {
		for c := 0; c < len(m.Tiles); c++ {
			assert.NotNil(t, m.Tiles[r][c])
		}
	}
}

func TestPlayerSpawning(t *testing.T) {
	m := NewMap(1)
	p := &Player{"Parth", 0, 0, false, "none", nil}
	x, y := m.SpawnNewPlayer(p)
	assert.NotNil(t, m.Tiles[y][x].Snake)

	found := false
	for r := range m.Tiles {
		for c := range m.Tiles[0] {
			if m.Tiles[r][c].Snake != nil && m.Tiles[r][c].Snake.Player != nil {
				found = true
			}
		}
	}

	assert.True(t, found, "Food should not have been found on this map")

	found = false

	for r := range m.Tiles {
		for c := range m.Tiles[0] {
			if m.Tiles[r][c].Food != nil {
				found = true
			}
		}
	}

	assert.False(t, found, "Food should not have been found on this map")
	assert.NotNil(t, p.Snake, "Snake not created properly")
}

func TestPlayerMovement0ToEdge(t *testing.T) {
	m := NewMap(1)
	p := &Player{"Parth", 0, 0, false, "none", nil}
	x, y := m.SpawnNewPlayer(p)

	p.CurrentDirection = 0
	updatesLeft := len(m.Tiles) - y - 1

	for i := 0; i < updatesLeft; i++ {
		m.Update()
	}

	assert.Equal(t, x, p.Snake.Head.X, "Snake did not move in the right direction")
	assert.Equal(t, y+updatesLeft, p.Snake.Head.Y, "Snake did not move in the right direction")

	m.Update()
	assert.Equal(t, 0, p.Snake.Length)
	assert.Nil(t, p.Snake.Head)
	assert.Equal(t, p.Snake, m.Losers[p])
}

func TestPlayerMovement1ToEdge(t *testing.T) {
	m := NewMap(1)
	p := &Player{"Parth", 0, 0, false, "none", nil}
	x, y := m.SpawnNewPlayer(p)

	p.CurrentDirection = 1
	updatesLeft := len(m.Tiles[0]) - x - 1

	for i := 0; i < updatesLeft; i++ {
		m.Update()
	}

	assert.Equal(t, x+updatesLeft, p.Snake.Head.X, "Snake did not move in the right direction")
	assert.Equal(t, y, p.Snake.Head.Y, "Snake did not move in the right direction")

	m.Update()
	assert.Equal(t, 0, p.Snake.Length)
	assert.Nil(t, p.Snake.Head)
	assert.Equal(t, p.Snake, m.Losers[p])
}

func TestPlayerMovement2ToEdge(t *testing.T) {
	m := NewMap(1)
	p := &Player{"Parth", 0, 0, false, "none", nil}
	x, y := m.SpawnNewPlayer(p)

	p.CurrentDirection = 2
	updatesLeft := y

	for i := 0; i < updatesLeft; i++ {
		m.Update()
	}

	assert.Equal(t, x, p.Snake.Head.X, "Snake did not move in the right direction")
	assert.Equal(t, 0, p.Snake.Head.Y, "Snake did not move in the right direction")

	m.Update()
	assert.Equal(t, 0, p.Snake.Length)
	assert.Nil(t, p.Snake.Head)
	assert.Equal(t, p.Snake, m.Losers[p])
}

func TestPlayerMovement3ToEdge(t *testing.T) {
	m := NewMap(1)
	p := &Player{"Parth", 0, 0, false, "none", nil}
	x, y := m.SpawnNewPlayer(p)

	p.CurrentDirection = 3
	updatesLeft := x

	for i := 0; i < updatesLeft; i++ {
		m.Update()
	}

	assert.Equal(t, 0, p.Snake.Head.X, "Snake did not move in the right direction")
	assert.Equal(t, y, p.Snake.Head.Y, "Snake did not move in the right direction")
	m.Update()
	assert.Equal(t, 0, p.Snake.Length)
	assert.Nil(t, p.Snake.Head)
	assert.Equal(t, p.Snake, m.Losers[p])
}

func TestFoodDynamics(t *testing.T) {
	m := NewMap(1)
	p := &Player{"Parth", 0, 0, false, "none", nil}
	m.SpawnNewPlayerAt(p, 3, 3)
	m.AddFood(&Food{3, 5})
	m.AddFood(&Food{5, 5})
	p.CurrentDirection = 0

	assert.NotNil(t, m.Tiles[3][3].Snake)
	assert.NotNil(t, m.Tiles[5][3].Food)
	assert.NotNil(t, m.Tiles[5][5].Food)

	m.Update()
	m.Update()

	assert.Equal(t, 2, p.Snake.Length)

	p.CurrentDirection = 1

	m.Update()
	assert.Equal(t, 2, p.Snake.Length)

	m.Update()
	assert.Equal(t, 3, p.Snake.Length)

	for i := 0; i < 5; i++ {
		m.Update()
	}

	assert.Nil(t, m.Tiles[5][9].Snake)
	assert.Nil(t, m.Tiles[5][8].Snake)
	assert.Nil(t, m.Tiles[5][7].Snake)

	assert.NotNil(t, m.Tiles[5][9].Food)
	assert.NotNil(t, m.Tiles[5][8].Food)
	assert.NotNil(t, m.Tiles[5][7].Food)
}

func TestMultiPlayerDynamics(t *testing.T) {
}
