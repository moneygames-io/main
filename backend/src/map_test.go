package main

import (
	"testing"
	"fmt"
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
	m := NewMap(1)
	printMap(m)
}

func TestPlayerSpawning(t *testing.T) {
}

func TestPlayerMovement(t *testing.T) {
}

func TestMultiPlayerDynamics(t *testing.T) {
}
