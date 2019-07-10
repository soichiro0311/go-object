package domain

type Game struct {
	player   *Player
	dealer   *Dealer
	betMoney BetMoney
}

func (g *Game) Initilize(betMoney BetMoney) {
	deck := NewDeck()
	playerHand := NewHand(deck.draw(), deck.draw())
	g.player = NewPlayer(playerHand)
	dealerHand := NewHand(deck.draw(), deck.draw())
	g.dealer = NewDealer(dealerHand)
	g.betMoney = betMoney
}

func (g *Game) ReferrencePlayerHand() []Card {
	return g.player.hand.cards
}

func (g *Game) ReferrenceDealerHand() []Card {
	return g.dealer.hand.cards
}

func (g *Game) ReferrenceBetMoney() BetMoney {
	return g.betMoney
}
