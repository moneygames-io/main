package main

type Player struct {
	Name             string
	CurrentDirection int
	CurrentSprint    bool
	Token		     string
	Snake            *Snake
}
