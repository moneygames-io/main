package main

type Snake struct {
	Player string
	Length int
	CurrentDirection int
	Head   *SnakeNode
	Tail   *SnakeNode
	Events MapEvent
}

func NewSnake(x int, y int, events MapEvent) *Snake {
	s := new(Snake)
	s.Head = &SnakeNode{s, x, y, nil, nil}
	s.Tail = s.Head
	s.Length = 1
	s.Events = events

	s.Events.AddSnakeNode(s.Head)
	return s
}

func (snake *Snake) Move(direction int) {
	if math.Abs(direction-snake.CurrentDirection) != 2 {
		snake.CurrentDirection = direction
	}
	snake.Grow()
	snake.Shorten(1)
}

func (snake *Snake) Sprint(direction int) {
	if math.Abs(direction-snake.CurrentDirection) != 2 {
		snake.CurrentDirection = direction
	}
	snake.grow()
	oldTail := snake.Shorten(2)
	snake.Events.AddFood(Food(oldTail.X, oldTail.Y))
}

func (snake *Snake) Grow() {
	oldHead := snake.Head
	dx, dy = directionToDxDy(snake.CurrentDirection)

	newHead := new(SnakeNode)
	newHead.Snake = oldHead.Snake
	newHead.X = dx + oldHead.X
	newHead.Y = dy + oldHead.Y
	newHead.Next = oldHead
	newHead.Prev = nil

	oldHead.Prev = newHead

	snake.Head = newHead
	snake.Length = snake.Length+1

	snake.Events.AddSnakeNode(newHead)
}

func (snake *Snake) Shorten(howMuch int) SnakeNode* {
	oldTail := snake.Tail
	newTail := snake.Tail.Prev

	oldTail.Prev = nil
	newTail.Next = nil

	snake.Tail = newTail
	snake.Length = snake.Length - 1

	snake.Events.RemoveSnakeNode(oldTail)

	if howMuch > 1 {
		return snake.Shorten(howMuch - 1)
	} else {
		return oldTail
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
