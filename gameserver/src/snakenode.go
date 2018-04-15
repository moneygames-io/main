package main

type SnakeNode struct {
	Snake *Snake
	X     int
	Y     int
	Next  *SnakeNode
	Prev  *SnakeNode
}
