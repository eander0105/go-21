package player_test

import (
	"testing"

	. "github.com/eander0105/go-21/card"
	. "github.com/eander0105/go-21/player"
)

// Utils
func flipAllCards(h *Hand) {
	for i := range h.Cards {
		h.Cards[i].Flip()
	}
}

// Tests
func TestBasicHandValues(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			NewCard(Hearts, King),
			NewCard(Diamonds, King),
		},
	}
	flipAllCards(&hand)

	if hand.Value() != 20 {
		t.Errorf("Expected 20, got %d", hand.Value())
	}

	hand = Hand{
		Cards: []Card{
			NewCard(Hearts, Two),   // 2
			NewCard(Hearts, Three), // 3
			NewCard(Hearts, Four),  // 4
			NewCard(Hearts, Five),  // 5
			NewCard(Hearts, Six),   // 6
			NewCard(Hearts, Seven), // 7
			NewCard(Hearts, Eight), // 8
			NewCard(Hearts, Nine),  // 9
			NewCard(Hearts, Ten),   // 10
			NewCard(Hearts, Jack),  // 10
			NewCard(Hearts, Queen), // 10
			NewCard(Hearts, King),  // 10 = 84
		},
	}
	flipAllCards(&hand)

	if hand.Value() != 84 {
		t.Errorf("Expected 74, got %d", hand.Value())
	}
}

func TestIsSoft(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			NewCard(Hearts, Ace),
			NewCard(Diamonds, King),
		},
	}
	flipAllCards(&hand)

	if !hand.IsSoft() {
		t.Errorf("Expected true, got false")
	}

	hand = Hand{
		Cards: []Card{
			NewCard(Hearts, Two),   // 2
			NewCard(Hearts, Three), // 3
			NewCard(Hearts, Four),  // 4
			NewCard(Hearts, Five),  // 5
			NewCard(Hearts, Six),   // 6
			NewCard(Hearts, Seven), // 7
			NewCard(Hearts, Eight), // 8
			NewCard(Hearts, Nine),  // 9
			NewCard(Hearts, Ten),   // 10
			NewCard(Hearts, Jack),  // 10
			NewCard(Hearts, Queen), // 10
			NewCard(Hearts, King),  // 10
			NewCard(Hearts, Ace),   // 11
		},
	}
	flipAllCards(&hand)

	if hand.IsSoft() {
		t.Errorf("Expected false, got true")
	}
}

func TestFlipAllCards(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			NewCard(Hearts, Ace),
			NewCard(Diamonds, King),
			NewCard(Clubs, Queen),
		},
	}

	for _, card := range hand.Cards {
		if card.FaceUp {
			t.Errorf("Expected card to be face down, but it's face up")
		}
	}

	flipAllCards(&hand)

	for _, card := range hand.Cards {
		if !card.FaceUp {
			t.Errorf("Expected card to be face up, but it's face down")
		}
	}
}

func TestCalculateHandValue(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			NewCard(Hearts, Ace),
			NewCard(Diamonds, King),
		},
	}
	flipAllCards(&hand)

	value := hand.Value()
	if value != 21 {
		t.Errorf("Expected hand value to be 21, but it's %d", value)
	}
	if !hand.IsSoft() {
		t.Errorf("Expected hand to be soft, but it's not")
	}
}

func TestCalculateHandValueMultipleCards(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			NewCard(Hearts, Ace),
			NewCard(Diamonds, King),
			NewCard(Clubs, Queen),
			NewCard(Spades, Jack),
		},
	}
	flipAllCards(&hand)

	value := hand.Value()
	if value != 31 {
		t.Errorf("Expected hand value to be 31, but it's %d", value)
	}
	if hand.IsSoft() {
		t.Errorf("Expected hand to be hard, but it's soft")
	}
}

func TestIsBusted(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			NewCard(Hearts, Ace),
			NewCard(Diamonds, King),
		},
	}
	flipAllCards(&hand)

	if hand.IsBusted() {
		t.Errorf("Expected false, got true")
	}

	hand = Hand{
		Cards: []Card{
			NewCard(Hearts, King),
			NewCard(Diamonds, King),
			NewCard(Clubs, King),
		},
	}
	flipAllCards(&hand)

}
