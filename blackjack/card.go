package blackjack

import "math/rand"

type Card struct {
	Suit   string
	Value  string
	FaceUp bool
}

type Deck []Card

func NewCard(suit, value string) Card {
	return Card{
		Suit:   suit,
		Value:  value,
		FaceUp: false,
	}
}

func (c Card) String() string {
	if c.FaceUp {
		return c.Value + " of " + c.Suit
	}
	return "Face Down"
}

func (d *Deck) DrawTopCard(faceUp bool) Card {
	card := (*d)[0]

	if faceUp {
		card.FaceUp = true
	}

	*d = (*d)[1:]
	return card
}

func (d Deck) Shuffle() {
	for i := range d {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

func NewDeck(shuffled bool, noDecks int) Deck {
	if noDecks < 1 {
		noDecks = 1
	}
	d := Deck{}

	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for i := 0; i < noDecks; i++ {
		for _, suit := range suits {
			for _, value := range values {
				d = append(d, NewCard(suit, value))
			}
		}
	}

	if shuffled {
		d.Shuffle()
	}

	return d
}
