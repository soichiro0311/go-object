package interfaces

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"../domain"
)

func GameStartPrinter(game *domain.Game) {
	fmt.Println("賭け金：【" + game.ReferrenceBetMoney().ToString() + "】")
	playerHand := game.ReferrencePlayerHand()
	fmt.Println("Playerのカード：【" + playerHand[0].ToString() + "】,【" + playerHand[1].ToString() + "】")
	dealerHand := game.ReferrenceDealerHand()
	fmt.Println("Dealerのカード：【" + dealerHand[0].ToString() + "】,【?】")
}
func BetMoneyRequire() string {
	fmt.Println("賭け金を入力してください")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	betMoneyString := scanner.Text()
	betMoneyString = strings.TrimSpace(betMoneyString)

	return betMoneyString
}
