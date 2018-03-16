package main

type Snake struct {
	Player string
	Length int
	Head *SnakeNode
	Tail *SnakeNode
	Map *Map
}

type SnakeNode struct {
	Snake *Snake
	X int
	Y int
	Next *SnakeNode
	Prev *SnakeNode
}

func (snake *Snake) NewSnake(x int, y int, world *Map) *Snake {
	s := new(Snake)
	s.Head = &SnakeNode{s, x, y, nil, nil}
	s.Tail = s.Head
	s.Length = 1
	s.Map = world
	return s
}

func (snake *Snake) move(direction int) {

}
