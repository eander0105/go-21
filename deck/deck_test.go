package deck_test

import (
	"testing"

	// . "github.com/eander0105/go-21/card"
	. "github.com/eander0105/go-21/deck"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck(1, true)
	if len(d) != 52 {
		t.Errorf("Expected 52 cards, got %d", len(d))
	}

	d = NewDeck(2, true)
	if len(d) != 104 {
		t.Errorf("Expected 104 cards, got %d", len(d))
	}

	d = NewDeck(0, true)
	if len(d) != 52 {
		t.Errorf("Expected 52 cards even though less than 1 deck was inputed, got %d", len(d))
	}
}

func TestShuffle(t *testing.T) {
	d := NewDeck(1, false)
	firstCard := d[0]
	d.Shuffle()
	if d[0] == firstCard {
		t.Errorf("Expected %s to be different from %s", d[0], firstCard)
	}
}

func TestDrawTopCard(t *testing.T) {
	d := NewDeck(1, false)
	c := d.DrawTopCard(true)

	if c.String() != "Ace of Hearts" {
		t.Errorf("Expected Ace of Hearts, got %s", c.String())
	}

	if len(d) != 51 {
		t.Errorf("Expected 51 cards, got %d", len(d))
	}

	c = d.DrawTopCard(false)

	if c.String() != "Face Down" {
		t.Errorf("Expected Face down, got %s", c.String())
	}

	if len(d) != 50 {
		t.Errorf("Expected 50 cards, got %d", len(d))
	}
}
