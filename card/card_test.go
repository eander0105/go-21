package card_test

import (
	"testing"

	. "github.com/eander0105/go-21/card"
)

func TestCard(t *testing.T) {
	c := NewCard(Hearts, Ace)
	if c.Value != Ace {
		t.Errorf("Expected Ace, got %s", c.Value)
	}

	if c.Suit != Hearts {
		t.Errorf("Expected Hearts, got %s", c.Suit)
	}

	if c.FaceUp {
		t.Errorf("Face up was true, expected false")
	}

	if c.String() != "Face Down" {
		t.Errorf("Expected Face down, got %s", c.String())
	}

	c.FaceUp = true

	if !c.FaceUp {
		t.Errorf("Face up was false, expected true")
	}

	if c.String() != "Ace of Hearts" {
		t.Errorf("Expected Ace of Hearts, got %s", c.String())
	}
}
