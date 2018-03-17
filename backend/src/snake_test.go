package main

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

type DummyMap struct {
	AddSN int
	RemoveSN int
	AddF int
	RemoveF int
}

func (dm DummyMap) AddSnakeNode(sn *SnakeNode) {
	dm.AddSN = dm.AddSN+1
}

func (dm DummyMap) RemoveSnakeNode (sn *SnakeNode) {
	dm.RemoveSN++
}

func (dm DummyMap) AddFood (sn *Food) {
	dm.AddF++
}

func (dm DummyMap) RemoveFood (sn *Food) {
	dm.RemoveF++
}

func TestNewSnake(t *testing.T) {
	dummyMap := DummyMap{0, 0, 0, 0}
	testSnake := NewSnake(0, 0, dummyMap)
	assert.Equal(t, testSnake.Length, 1, "new snake length should be 1")
	assert.Equal(t, testSnake.CurrentDirection, 0, "default current direction should be 0")
	fmt.Println(dummyMap)
	assert.Equal(t, 1, dummyMap.AddSN);
}
