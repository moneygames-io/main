package main

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func printMap(m *Map) {
	for r := range m.Tiles {
		for c := range m.Tiles[0]{
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
	assert.NotNil(t, m.Tiles[0])
	assert.NotNil(t, m.Players)
}

func TestPlayerSpawning(t *testing.T) {
	m := NewMap(1)
	p := &Player{"Parth", 0, 0, false, "none", nil}
	m.SpawnNewPlayer(p)

	found := false
	for r := range m.Tiles {
		for c := range m.Tiles[0]{
			if m.Tiles[r][c].Snake != nil && m.Tiles[r][c].Snake.Player != nil {
				found = true
			}
		}
	}

	assert.True(t, found, "Food should not have been found on this map")

	found = false

	for r := range m.Tiles {
		for c := range m.Tiles[0]{
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
		m.update()
	}

	assert.Equal(t, x, p.Snake.Head.X, "Snake did not move in the right direction")
	assert.Equal(t, y + updatesLeft, p.Snake.Head.Y, "Snake did not move in the right direction")
}

func TestPlayerMovement1ToEdge(t *testing.T) {
	m := NewMap(1)
	p := &Player{"Parth", 0, 0, false, "none", nil}
	x, y := m.SpawnNewPlayer(p)

	p.CurrentDirection = 1
	updatesLeft := len(m.Tiles[0]) - x - 1

	for i := 0; i < updatesLeft; i++ {
		m.update()
	}

	assert.Equal(t, x + updatesLeft, p.Snake.Head.X, "Snake did not move in the right direction")
	assert.Equal(t, y, p.Snake.Head.Y, "Snake did not move in the right direction")
}

func TestPlayerMovement2ToEdge(t *testing.T) {
	m := NewMap(1)
	p := &Player{"Parth", 0, 0, false, "none", nil}
	x, y := m.SpawnNewPlayer(p)

	p.CurrentDirection = 2
	updatesLeft := y

	for i := 0; i < updatesLeft; i++ {
		m.update()
	}

	assert.Equal(t, x, p.Snake.Head.X, "Snake did not move in the right direction")
	assert.Equal(t, 0, p.Snake.Head.Y, "Snake did not move in the right direction")
}

func TestPlayerMovement3ToEdge(t *testing.T) {
	m := NewMap(1)
	p := &Player{"Parth", 0, 0, false, "none", nil}
	x, y := m.SpawnNewPlayer(p)

	p.CurrentDirection = 3
	updatesLeft := x

	for i := 0; i < updatesLeft; i++ {
		m.update()
	}

	assert.Equal(t, 0, p.Snake.Head.X, "Snake did not move in the right direction")
	assert.Equal(t, y, p.Snake.Head.Y, "Snake did not move in the right direction")
}

func TestMultiPlayerDynamics(t *testing.T) {
}
