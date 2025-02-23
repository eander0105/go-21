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

	if !hand.IsBusted() {
		t.Errorf("Expected true, got false")
	}
}

func TestIsBlackJack(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			NewCard(Hearts, Ace),
			NewCard(Diamonds, King),
		},
	}
	flipAllCards(&hand)

	if !hand.IsBlackJack() {
		t.Errorf("Expected true, got false")
	}

	hand = Hand{
		Cards: []Card{
			NewCard(Hearts, King),
			NewCard(Diamonds, King),
			NewCard(Clubs, King),
		},
	}
	flipAllCards(&hand)

	if hand.IsBlackJack() {
		t.Errorf("Expected false, got true")
	}
}

func TestPartialyFaceUpHand(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			NewCard(Hearts, Ace),
			NewCard(Diamonds, King),
			NewCard(Clubs, Queen),
		},
	}

	for i := 0; i < len(hand.Cards)-1; i++ {
		hand.Cards[i].Flip()
	}

	value := hand.Value()

	if value != 21 {
		t.Errorf("Expected hand value to be 21, but it's %d", value)
	}

}

func TestAddCard(t *testing.T) {
	hand := Hand{
		Cards: []Card{},
	}

	hand.Hit(NewCard(Hearts, Ace))
	hand.Cards[0].Flip()

	if len(hand.Cards) != 1 {
		t.Errorf("Expected hand to have 1 card, but it has %d", len(hand.Cards))
	}

	if hand.Value() != 11 {
		t.Errorf("Expected hand value to be 11, but it's %d", hand.Value())
	}

	if !hand.IsSoft() {
		t.Errorf("Expected hand to be soft, but it's not")
	}

	hand.Hit(NewCard(Hearts, King))
	hand.Cards[1].Flip()

	if hand.Value() != 21 {
		t.Errorf("Expected hand value to be 21, but it's %d", hand.Value())
	}

	if !hand.IsSoft() {
		t.Errorf("Expected hand to be soft, but it's hard")
	}

	if hand.IsBusted() {
		t.Errorf("Expected false, got true")
	}

	if !hand.IsBlackJack() {
		t.Errorf("Expected true, got false")
	}

	hand.Hit(NewCard(Hearts, Queen))
	hand.Cards[2].Flip()

	if hand.Value() != 21 {
		t.Errorf("Expected hand value to be 21, but it's %d", hand.Value())
	}

	if hand.IsSoft() {
		t.Errorf("Expected hand to be hard, but it's soft")
	}

	if hand.IsBusted() {
		t.Errorf("Expected false, got true")
	}

	if hand.IsBlackJack() {
		t.Errorf("Expected false, got true")
	}

	hand.Hit(NewCard(Hearts, Jack))
	hand.Cards[3].Flip()

	if !hand.IsBusted() {
		t.Errorf("Expected true, got false")
	}
}

func TestString(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			NewCard(Hearts, Ace),
			NewCard(Diamonds, King),
		},
	}
	flipAllCards(&hand)

	expected := "Ace of Hearts, King of Diamonds. 21 (Soft)"
	if hand.String() != expected {
		t.Errorf("Expected %s, got %s", expected, hand.String())
	}

	hand = Hand{
		Cards: []Card{
			NewCard(Hearts, King),
			NewCard(Diamonds, King),
			NewCard(Clubs, King),
		},
	}
	flipAllCards(&hand)

	expected = "King of Hearts, King of Diamonds, King of Clubs. 30"
	if hand.String() != expected {
		t.Errorf("Expected %s, got %s", expected, hand.String())
	}

	hand = Hand{
		Cards: []Card{
			NewCard(Hearts, King),
			NewCard(Diamonds, Ace),
		},
	}
	hand.Cards[0].Flip()

	expected = "King of Hearts, Face Down. 10"
	if hand.String() != expected {
		t.Errorf("Expected %s, got %s", expected, hand.String())
	}
}

func TestIsLocked(t *testing.T) {
	hand := Hand{
		Cards: []Card{
			NewCard(Hearts, Ace),
			NewCard(Diamonds, King),
		},
	}
	flipAllCards(&hand)

	if hand.Locked {
		t.Errorf("Expected false, got true")
	}

	hand.Stand()

	if !hand.Locked {
		t.Errorf("Expected true, got false")
	}
}

func TestHitAndStand(t *testing.T) {
	hand := Hand{
		Cards: []Card{},
	}

	hand.Hit(NewCard(Hearts, Ace))
	hand.Hit(NewCard(Diamonds, King))
	flipAllCards(&hand)

	if hand.Value() != 21 {
		t.Errorf("Expected hand value to be 21, but it's %d", hand.Value())
	}

	if hand.Locked {
		t.Errorf("Expected false, got true")
	}

	hand.Stand()

	if !hand.Locked {
		t.Errorf("Expected true, got false")
	}

	hand.Hit(NewCard(Hearts, King))

	if len(hand.Cards) != 2 {
		t.Errorf("Expected hand to have 2 cards, but it has %d", len(hand.Cards))
	}

	if hand.Value() != 21 {
		t.Errorf("Expected hand value to be 21, but it's %d", hand.Value())
	}

	if !hand.Locked {
		t.Errorf("Expected false, got true")
	}
}
