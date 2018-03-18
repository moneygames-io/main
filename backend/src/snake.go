package main

import (
	"math"
)

type Snake struct {
	Player           string
	Length           int
	CurrentDirection int
	Head             *SnakeNode
	Tail             *SnakeNode
	Events           MapEvent
}

func NewSnake(x int, y int, events MapEvent) *Snake {
	s := new(Snake)
	s.Head = &SnakeNode{s, x, y, nil, nil}
	s.Tail = s.Head
	s.Length = 1
	s.CurrentDirection = 0
	s.Events = events

	s.Events.SnakeCreated(s)
	return s
}

func (snake *Snake) Move(direction int) {
	if snake.Length == 1 || math.Abs(float64(direction-snake.CurrentDirection)) != 2 {
		snake.CurrentDirection = direction
	}
	status := snake.GrowHead()

	switch status {
	case 0: // Empty Cell
		snake.ShortenTail(1)
		break
	case 1: // Food
		break
	case 2:
		snake.ShortenTail(1)
		snake.Dead()
	}
}

func (snake *Snake) Sprint(direction int) {
	if math.Abs(float64(direction-snake.CurrentDirection)) != 2 {
		snake.CurrentDirection = direction
	}

	status := snake.GrowHead()

	switch status {
	case 0:
		oldTail := snake.ShortenTail(2)
		snake.Events.AddFood(&Food{oldTail.X, oldTail.Y})
	case 1:
		oldTail := snake.ShortenTail(1)
		snake.Events.AddFood(&Food{oldTail.X, oldTail.Y})
	case 2:
		snake.Dead()
	}

}

func (snake *Snake) GrowHead() int {
	oldHead := snake.Head
	dx, dy := directionToDxDy(snake.CurrentDirection)

	newHead := new(SnakeNode)
	newHead.Snake = oldHead.Snake
	newHead.X = dx + oldHead.X
	newHead.Y = dy + oldHead.Y
	newHead.Next = oldHead
	newHead.Prev = nil

	oldHead.Prev = newHead

	snake.Head = newHead
	snake.Length = snake.Length + 1

	return snake.Events.HeadMoved(snake)
}

func (snake *Snake) ShortenTail(howMuch int) *SnakeNode {
	if snake.Length == 1 { // no tail left
		return nil
	}
	oldTail := snake.Tail
	newTail := snake.Tail.Prev

	newTail.Next = nil

	snake.Tail = newTail
	snake.Length = snake.Length - 1

	snake.Events.RemoveTailNode(oldTail)

	if howMuch > 1 { // more tail to get rid off
		return snake.ShortenTail(howMuch - 1)
	} else {
		return oldTail
	}
}

func (snake *Snake) Dead() {
	sn := snake.ShortenTail(1)
	if sn != nil {
		snake.Events.AddFood(&Food{sn.X, sn.Y})
		snake.Dead()
	} else {
		snake.Length = 0
		snake.Events.RemoveTailNode(snake.Head)
		snake.Events.SnakeRemoved(snake)
		snake.Events.AddFood(&Food{snake.Head.X, snake.Head.Y})
		snake.Tail = nil
		snake.Head = nil
	}
}

func directionToDxDy(direction int) (int, int) {
	switch direction {
	case 0:
		return 0, 1
	case 1:
		return 1, 0
	case 2:
		return 0, -1
	case 3:
		return -1, 0
	default:
		return 0, 0
	}
}
