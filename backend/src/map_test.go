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
}

func TestNewMap(t *testing.T) {
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

}

func TestPlayerMovement(t *testing.T) {
}

func TestMultiPlayerDynamics(t *testing.T) {
}
