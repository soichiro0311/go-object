package domain

type Game struct {
	player *Player
	dealer *Dealer
}

func (g *Game) Initilize() {
	deck := NewDeck()
	playerHand := NewHand(deck.draw(), deck.draw())
	g.player = NewPlayer(playerHand)
	dealerHand := NewHand(deck.draw(), deck.draw())
	g.dealer = NewDealer(dealerHand)
}

func (g *Game) ReferrencePlayerHand() []Card {
	return g.player.hand.cards
}

func (g *Game) ReferrenceDealerHand() []Card {
	return g.dealer.hand.cards
}
