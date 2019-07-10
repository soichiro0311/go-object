package domain

type Hand struct {
	cards []Card
}

func NewHand(card1 Card, card2 Card) Hand {
	hand := new(Hand)
	hand.cards = append(hand.cards, card1, card2)
	return *hand
}
