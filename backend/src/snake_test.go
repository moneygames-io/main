package main

import (
	"testing"
)

func TestNewSnake(t *testing.T) {
	testSnake := NewSnake(0, 0, nil)
	if testSnake.Map != nil {
		t.Error("expected", nil, "got", testSnake.Map)
	}
}
