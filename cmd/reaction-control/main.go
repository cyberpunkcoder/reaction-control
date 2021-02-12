package main

import (
	"log"

	rice "github.com/GeertJohan/go.rice"
	"github.com/cyberpunkcoder/reaction-control/internal/game"
)

var g game.Game

func main() {
	// must reference rice lib usage here otherwise generator cannot find usages
	assetsBox, err := rice.FindBox("../../assets")
	if err != nil {
		log.Fatalf("Unable to initialise assets rice box...: %v\n", err)
	}

	g = game.Game{}
	g.Start(assetsBox)
}
