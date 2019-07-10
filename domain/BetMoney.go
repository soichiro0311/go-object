package domain

import (
	"errors"
	"strconv"
)

type BetMoney struct {
	amount int
}

func NewBetMoney(amount int) (BetMoney, error) {
	money := new(BetMoney)
	money.amount = amount
	if err := money.validateAmount(); err != nil {
		return *money, err
	}
	return *money, nil
}

func (b BetMoney) validateAmount() error {
	if b.amount > 0 {
		return nil
	}
	return errors.New("賭け金は0円以下は認められません")
}

func (b BetMoney) ToString() string {
	return strconv.Itoa(b.amount)
}
