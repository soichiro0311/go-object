package domain

import (
	"strconv"

	"./enum"
)

type Card struct {
	number int
	suit   enum.SuitEnum
}

func NewCard(number int, suit enum.SuitEnum) Card {
	card := new(Card)
	card.number = number
	card.suit = suit
	return *card
}

func (card Card) ToString() string {
	return "数字:" + strconv.Itoa(card.number) + " 絵柄：" + card.suit.ToString()
}
