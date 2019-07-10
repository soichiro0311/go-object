package main

import (
	"./domain"
	"./interfaces"
)

func main() {
	game := new(domain.Game)
	game.Initilize()
	interfaces.GameStartPrinter(game)
}
