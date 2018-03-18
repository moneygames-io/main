package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type DummyMap struct {
	Created        int
	AddSN          int
	RemoveSN       int
	AddF           int
	RemoveF        int
	IntendedReturn int
}

func (dm *DummyMap) SnakeCreated(sn *Snake) {
	dm.Created++
}

func (dm *DummyMap) HeadMoved(sn *Snake) int {
	dm.AddSN++
	return 0
}

func (dm *DummyMap) RemoveTailNode(sn *SnakeNode) {
	dm.RemoveSN++
}

func (dm *DummyMap) AddFood(sn *Food) {
	dm.AddF++
}

func (dm *DummyMap) RemoveFood(sn *Food) {
	dm.RemoveF++
}

func TestNewSnake(t *testing.T) {
	dummyMap := &DummyMap{IntendedReturn: 0}
	testSnake := NewSnake(0, 0, dummyMap)

	assert.Equal(t, testSnake.Length, 1, "new snake length should be 1")
	assert.Equal(t, testSnake.CurrentDirection, 0, "default current direction should be 0")
	assert.Equal(t, 1, dummyMap.Created)
}

func TestSnakeMoving(t *testing.T) {
}
