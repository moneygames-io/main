package main

type Player struct {
	Name string
	CurrentDirection int
	CurrentZoomLevel int
	CurrentSprint bool
	WalletAddress string
	Snake *Snake
}
