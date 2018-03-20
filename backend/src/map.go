package main

type Tile struct {
	Snake *Snake
	Depth int
}

type Map struct {
	Tiles   [][]Tile
	Players map[string]*Snake
}

type MapEvent interface {
	SnakeCreated(*Snake)
	HeadMoved(*Snake) int
	RemoveTailNode(*SnakeNode)
	AddFood(*Food)
	RemoveFood(*Food)
	SnakeRemoved(*Snake)
}

func NewMap(players int) *Map {
	newMap := new(Map)
	newMap.Tiles = [players * 100][players * 100]Tile{}
	for rows := 0; rows < len(newMap.Tiles); rows++ {
		for cols := 0; cols < len(newMap.Tiles[0]); cols++ {

		}
	}

	return newMap
}

func (m *Map) SpawnNewPlayer(player string) bool {

}

func (m *Map) SnakeCreated(*Snake) {

}

func (m *Map) HeadMoved(*Snake) int {

}

func (m *Map) RemoveTailNode(*SnakeNode) {

}

func (m *Map) AddFood(*Food) {

}

func (m *Map) RemoveFood(*Food) {

}

func (m *Map) SnakeRemoved(*Snake) {

}

func (m *Map) tick() {

}
