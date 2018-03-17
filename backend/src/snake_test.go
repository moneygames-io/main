package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type DummyMap struct {
}

func (dm DummyMap) AddSnakeNode(sn *SnakeNode) {
}

func (dm DummyMap) RemoveSnakeNode (sn *SnakeNode) {
}

func (dm DummyMap) AddFood (sn *Food) {
}

func (dm DummyMap) RemoveFood (sn *Food) {
}

func TestNewSnake(t *testing.T) {
	dummyMap := DummyMap{}
	testSnake := NewSnake(0, 0, dummyMap)
	assert.Equal(t, testSnake.Length, 1, "new snake length should be 1")
	assert.Equal(t, testSnake.CurrentDirection, 0, "default current direction should be 0")
}
