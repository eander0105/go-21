package player_test

import (
	"testing"

	c "github.com/eander0105/go-21/card"
	. "github.com/eander0105/go-21/player"
)

func TestBasicHandValues(t *testing.T) {
	hand := Hand{
		Cards: []c.Card{
			c.NewCard(c.Hearts, c.King),
			c.NewCard(c.Diamonds, c.King),
		},
	}

	if hand.Value() != 20 {
		t.Errorf("Expected 20, got %d", hand.Value())
	}

	hand = Hand{
		Cards: []c.Card{
			c.NewCard(c.Hearts, c.Two),   // 2
			c.NewCard(c.Hearts, c.Three), // 3
			c.NewCard(c.Hearts, c.Four),  // 4
			c.NewCard(c.Hearts, c.Five),  // 5
			c.NewCard(c.Hearts, c.Six),   // 6
			c.NewCard(c.Hearts, c.Seven), // 7
			c.NewCard(c.Hearts, c.Eight), // 8
			c.NewCard(c.Hearts, c.Nine),  // 9
			c.NewCard(c.Hearts, c.Ten),   // 10
			c.NewCard(c.Hearts, c.Jack),  // 10
			c.NewCard(c.Hearts, c.Queen), // 10
			c.NewCard(c.Hearts, c.King),  // 10 = 74
		},
	}

	if hand.Value() != 74 {
		t.Errorf("Expected 74, got %d", hand.Value())
	}
}
