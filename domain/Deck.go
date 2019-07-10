package domain

import (
	"math/rand"

	"./enum"
)

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	deck := new(Deck)
	cards := make([]Card, 0)
	for _, value := range enum.SuitEnumValues() {
		for i := 1; i < 14; i++ {
			cards = append(cards, NewCard(i, value))
		}
	}
	deck.cards = cards
	deck.shuffle()

	return deck
}

func (deck *Deck) shuffle() {
	n := len(deck.cards)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	}
}

func (deck *Deck) draw() Card {
	card := deck.cards[0]
	cards := make([]Card, 0)
	for i, v := range deck.cards {
		if i != 0 {
			cards = append(cards, v)
		}
	}
	deck.cards = cards
	return card
}
