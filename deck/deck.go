package deck

import (
	"math/rand"

	. "github.com/eander0105/go-21/card"
)

type Deck []Card

func NewDeck(noDecks int, shuffled bool) Deck {
	if noDecks < 1 {
		noDecks = 1
	}
	d := Deck{}

	suits := []Suit{Hearts, Diamonds, Clubs, Spades}
	values := []Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
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

func (d Deck) Shuffle() {
	for i := range d {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

func (d *Deck) DrawTopCard(faceUp bool) Card {
	card := (*d)[0]

	if faceUp {
		card.FaceUp = true
	}

	*d = (*d)[1:]
	return card
}
