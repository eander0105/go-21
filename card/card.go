package card

type Card struct {
	Suit   Suit
	Value  Value
	FaceUp bool
}

func NewCard(suit Suit, value Value) Card {
	return Card{
		Suit:   suit,
		Value:  value,
		FaceUp: false,
	}
}

func (c Card) String() string {
	if c.FaceUp {
		return c.Value.String() + " of " + c.Suit.String()
	}
	return "Face Down"
}

func (c *Card) Flip() {
	c.FaceUp = !c.FaceUp
}
