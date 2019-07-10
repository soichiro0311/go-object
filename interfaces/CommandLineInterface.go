package interfaces

import (
	"fmt"

	"../domain"
)

func GameStartPrinter(game *domain.Game) {
	playerHand := game.ReferrencePlayerHand()
	fmt.Println("Playerのカード：【" + playerHand[0].ToString() + "】,【" + playerHand[1].ToString() + "】")
	dealerHand := game.ReferrenceDealerHand()
	fmt.Println("Dealerのカード：【" + dealerHand[0].ToString() + "】,【?】")
}
