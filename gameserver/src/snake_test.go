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
	SnakesR        int
	IntendedReturn int
}

func (dm *DummyMap) SnakeCreated(sn *Snake) {
	dm.Created++
}

func (dm *DummyMap) AddNode(sn *SnakeNode) int {
	dm.AddSN++
	return dm.IntendedReturn
}

func (dm *DummyMap) RemoveNode(row int, col int) {
	dm.RemoveSN++
}

func (dm *DummyMap) AddFood(sn *Food) {
	dm.AddF++
}

func (dm *DummyMap) RemoveFood(row int, col int) {
	dm.RemoveF++
}

func (dm *DummyMap) SnakeRemoved(sn *Snake) {
	dm.SnakesR++
}

func TestNewSnake(t *testing.T) {
	dummyMap := &DummyMap{IntendedReturn: 0}
	testSnake := NewSnake(0, 0, dummyMap)

	assert.Equal(t, testSnake.Length, 1, "new snake length should be 1")
	assert.Equal(t, testSnake.CurrentDirection, 0, "default current direction should be 0")
	assert.Equal(t, 1, dummyMap.Created)
	assert.Equal(t, 0, testSnake.Head.X)
	assert.Equal(t, 0, testSnake.Head.Y)
}

func TestSnakeMoving(t *testing.T) {
	dummyMap := &DummyMap{IntendedReturn: 1}
	testSnake := NewSnake(0, 0, dummyMap)

	assert.Equal(t, testSnake.Length, 1, "Initial starting size should be 1")

	assert.Equal(t, 0, testSnake.Head.X)
	assert.Equal(t, 0, testSnake.Head.Y)

	testSnake.Move(0)

	assert.Equal(t, 2, testSnake.Length)
	assert.Equal(t, 0, testSnake.Head.X)
	assert.Equal(t, 1, testSnake.Head.Y)
}

func TestSnakeMoving2(t *testing.T) {
	dummyMap := &DummyMap{IntendedReturn: 1}
	testSnake := NewSnake(0, 0, dummyMap)

	assert.Equal(t, testSnake.Length, 1, "Initial starting size should be 1")

	assert.Equal(t, 0, testSnake.Head.X)
	assert.Equal(t, 0, testSnake.Head.Y)

	testSnake.Move(1)
	testSnake.Move(1)
	testSnake.Move(1)
	testSnake.Move(1)

	assert.Equal(t, 5, testSnake.Length)
	assert.Equal(t, 4, testSnake.Head.X)
	assert.Equal(t, 0, testSnake.Head.Y)
}

func TestSnakeMoving3(t *testing.T) {
	dummyMap := &DummyMap{IntendedReturn: 1}
	testSnake := NewSnake(0, 0, dummyMap)

	assert.Equal(t, testSnake.Length, 1, "Initial starting size should be 1")

	assert.Equal(t, 0, testSnake.Head.X)
	assert.Equal(t, 0, testSnake.Head.Y)

	testSnake.Move(0)
	testSnake.Move(0)
	testSnake.Move(0)
	testSnake.Move(0)

	assert.Equal(t, 5, testSnake.Length)
	assert.Equal(t, 0, testSnake.Head.X)
	assert.Equal(t, 4, testSnake.Head.Y)
}

func TestSnakeMoving4(t *testing.T) {
	dummyMap := &DummyMap{IntendedReturn: 1}
	testSnake := NewSnake(0, 0, dummyMap)

	assert.Equal(t, testSnake.Length, 1, "Initial starting size should be 1")

	assert.Equal(t, 0, testSnake.Head.X)
	assert.Equal(t, 0, testSnake.Head.Y)

	testSnake.Move(2)
	testSnake.Move(2)
	testSnake.Move(2)
	testSnake.Move(2)

	assert.Equal(t, 5, testSnake.Length)
	assert.Equal(t, 0, testSnake.Head.X)
	assert.Equal(t, -4, testSnake.Head.Y)
}

func TestSnakeMoving5(t *testing.T) {
	dummyMap := &DummyMap{IntendedReturn: 1}
	testSnake := NewSnake(0, 0, dummyMap)

	assert.Equal(t, testSnake.Length, 1, "Initial starting size should be 1")

	assert.Equal(t, 0, testSnake.Head.X)
	assert.Equal(t, 0, testSnake.Head.Y)

	testSnake.Move(3)
	testSnake.Move(3)
	testSnake.Move(3)
	testSnake.Move(3)

	assert.Equal(t, 5, testSnake.Length)
	assert.Equal(t, -4, testSnake.Head.X)
	assert.Equal(t, 0, testSnake.Head.Y)
}

func TestSnakeDead(t *testing.T) {
	dummyMap := &DummyMap{IntendedReturn: 1}
	testSnake := NewSnake(0, 0, dummyMap)

	assert.Equal(t, testSnake.Length, 1, "Initial starting size should be 1")

	assert.Equal(t, 0, testSnake.Head.X)
	assert.Equal(t, 0, testSnake.Head.Y)

	testSnake.Move(0)
	testSnake.Move(0)
	testSnake.Move(0)
	testSnake.Move(0)

	dummyMap.IntendedReturn = 2

	testSnake.Move(0)

	assert.Equal(t, 5, dummyMap.AddF)
}

func TestPositions(t *testing.T) {
	dummyMap := &DummyMap{IntendedReturn: 1}
	testSnake := NewSnake(0, 0, dummyMap)

	assert.Equal(t, 0, testSnake.Head.X)
	assert.Equal(t, 0, testSnake.Head.Y)

	testSnake.Move(0)

	assert.Equal(t, 0, testSnake.Head.X)
	assert.Equal(t, 1, testSnake.Head.Y)

	assert.Equal(t, 0, testSnake.Head.Next.X)
	assert.Equal(t, 0, testSnake.Head.Next.Y)

	testSnake.Move(1)

	assert.Equal(t, 1, testSnake.Head.X)
	assert.Equal(t, 1, testSnake.Head.Y)

	assert.Equal(t, 0, testSnake.Head.Next.X)
	assert.Equal(t, 1, testSnake.Head.Next.Y)

	assert.Equal(t, 0, testSnake.Head.Next.Next.X)
	assert.Equal(t, 0, testSnake.Head.Next.Next.Y)

	testSnake.Move(2)

	assert.Equal(t, 1, testSnake.Head.X)
	assert.Equal(t, 0, testSnake.Head.Y)

	assert.Equal(t, 1, testSnake.Head.Next.X)
	assert.Equal(t, 1, testSnake.Head.Next.Y)

	assert.Equal(t, 0, testSnake.Head.Next.Next.X)
	assert.Equal(t, 1, testSnake.Head.Next.Next.Y)

	assert.Equal(t, 0, testSnake.Head.Next.Next.Next.X)
	assert.Equal(t, 0, testSnake.Head.Next.Next.Next.Y)

	testSnake.Move(2)

	assert.Equal(t, 1, testSnake.Head.X)
	assert.Equal(t, -1, testSnake.Head.Y)

	assert.Equal(t, 1, testSnake.Head.Next.X)
	assert.Equal(t, 0, testSnake.Head.Next.Y)

	assert.Equal(t, 1, testSnake.Head.Next.Next.X)
	assert.Equal(t, 1, testSnake.Head.Next.Next.Y)

	assert.Equal(t, 0, testSnake.Head.Next.Next.Next.X)
	assert.Equal(t, 1, testSnake.Head.Next.Next.Next.Y)

	assert.Equal(t, 0, testSnake.Head.Next.Next.Next.Next.X)
	assert.Equal(t, 0, testSnake.Head.Next.Next.Next.Next.Y)

	testSnake.Move(3)

	assert.Equal(t, 0, testSnake.Head.X)
	assert.Equal(t, -1, testSnake.Head.Y)

	assert.Equal(t, 1, testSnake.Head.Next.X)
	assert.Equal(t, -1, testSnake.Head.Next.Y)

	assert.Equal(t, 1, testSnake.Head.Next.Next.X)
	assert.Equal(t, 0, testSnake.Head.Next.Next.Y)

	assert.Equal(t, 1, testSnake.Head.Next.Next.Next.X)
	assert.Equal(t, 1, testSnake.Head.Next.Next.Next.Y)

	assert.Equal(t, 0, testSnake.Head.Next.Next.Next.Next.X)
	assert.Equal(t, 1, testSnake.Head.Next.Next.Next.Next.Y)

	assert.Equal(t, 0, testSnake.Head.Next.Next.Next.Next.Next.X)
	assert.Equal(t, 0, testSnake.Head.Next.Next.Next.Next.Next.Y)
}