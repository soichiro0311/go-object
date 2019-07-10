package domain

type Dealer struct {
	hand Hand
}

func NewDealer(hand Hand) *Dealer {
	dealer := new(Dealer)
	dealer.hand = hand
	return dealer
}
