package adapter

import (
	"log"
	"strconv"

	"../domain"
)

func ConvertToMoney(moneyString string) (domain.BetMoney, error) {
	amount, _ := strconv.Atoi(moneyString)
	money, err := domain.NewBetMoney(amount)
	if err != nil {
		log.Fatal(err)
		return money, err
	}
	return money, nil
}
