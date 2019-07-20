package domain

import (
	"errors"
	"strconv"

	"./enum"
)

type Card struct {
	number int
	suit   enum.SuitEnum
}

func NewCard(number int, suit enum.SuitEnum) (Card, error) {
	card := new(Card)
	card.number = number
	card.suit = suit
	err := card.validateNumber()
	return *card, err
}

func (card Card) ToString() string {
	return "数字:" + strconv.Itoa(card.number) + " 絵柄：" + card.suit.ToString()
}

func (card Card) validateNumber() error {
	if card.number < 0 || card.number > 13 {
		return errors.New("Card create failed. Reason is invalid number input.")
	}
	return nil
}
