package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewSnake(t *testing.T) {
	testSnake := NewSnake(0, 0, nil)
	assert.Equal(t, testSnake.Length, 1, "new snake length should be 1")
	assert.Equal(t, testSnake.CurrentDirection, 0, "default current direction should be 0")
}
