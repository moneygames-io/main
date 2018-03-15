type Snake struct {
	Player string
	Length int
	Head *SnakeNode
	Tail *SnakeNode
}

type SnakeNode {
	Snake Snake*
	X int
	Y int
	Next SnakeNode*
	Prev SnakeNode*
}

func (snake *Snake) NewSnake(x int, y int) *Snake {
	s := new(Snake)
	s.Head = &SnakeNode{&s, x, y, nil, nil}
	s.Tail = Head
	s.Length = 1
}

func (snake *Snake) move(direction int, [][][]int) {

}
