package enum

type SuitEnum int

const (
	SPADE = iota
	HEART
	DIAMOND
	CLUB
)

var suitEnumString = [4]string{"SPADE", "HEART", "DIAMOND", "CLUB"}

func SuitEnumValues() []SuitEnum {
	return []SuitEnum{SPADE, HEART, DIAMOND, CLUB}
}

func (suit SuitEnum) ToString() string {
	return suitEnumString[suit]
}
