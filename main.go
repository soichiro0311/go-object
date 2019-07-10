package main

import (
	"./adapter"
	"./domain"
	"./interfaces"
)

func main() {
	betMoney, err := adapter.ConvertToMoney(interfaces.BetMoneyRequire())
	if err != nil {
		panic(err)
	}
	game := new(domain.Game)
	game.Initilize(betMoney)
	interfaces.GameStartPrinter(game)
}
