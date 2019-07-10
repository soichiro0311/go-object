package domain

type Player struct {
	hand Hand
}

func NewPlayer(hand Hand) *Player {
	player := new(Player)
	player.hand = hand
	return player
}
