package main

type Player struct {
	Name string
	CurrentDirection int
	CurrentZoomLevel int
	WalletAddress string
	Snake *Snake
}
