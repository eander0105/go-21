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

var suitASCII = map[Suit]string{
	Hearts:   "♥",
	Diamonds: "♦",
	Clubs:    "♣",
	Spades:   "♠",
}

func (s Suit) String() string {
	return suitNames[s]
}

func (s Suit) ASCII() string {
	return suitASCII[s]
}
