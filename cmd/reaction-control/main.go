package main

import "github.com/cyberpunkprogrammer/reaction-control/internal/game"

var g game.Game

func main() {
	g = game.Game{}
	g.Start()
}
