package card

type Suit int

const (
	Hearts Suit = iota
	Diamonds
	Clubs
	Spades
)

var suitNames = map[Suit]string{
	Hearts:   "Hearts",
	Diamonds: "Diamonds",
	Clubs:    "Clubs",
	Spades:   "Spades",
}

func (s Suit) String() string {
	return suitNames[s]
}
