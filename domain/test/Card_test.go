package test

import (
	"testing"

	"../../domain"
	"../enum"
)

func Test_数字が0_13以内の場合にカードがエラーが発生せずに作成できること(t *testing.T) {
	card, err := domain.NewCard(5, enum.DIAMOND)
	if err != nil {
		t.Fatal(card.ToString())
	}
}

func Test_数字が0_13以外の場合にカードを作成しようとするとエラーが発生すること(t *testing.T) {
	card, err := domain.NewCard(14, enum.DIAMOND)
	if err == nil {
		t.Fatal(card.ToString())
	}
}
