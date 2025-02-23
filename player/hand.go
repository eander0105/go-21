package player

import (
	"slices"
	"strconv"
	"strings"

	c "github.com/eander0105/go-21/card"
)

type Hand struct {
	Cards  []c.Card
	Bet    float32
	Locked bool
}

func (h Hand) String() string {
	cards := strings.Join(func() []string {
		cStr := make([]string, len(h.Cards))
		for i, card := range h.Cards {
			cStr[i] = card.String()
		}
		return cStr
	}(), ", ")

	return cards + ". " + strconv.Itoa(h.Value()) + func() string {
		if h.IsSoft() {
			return " (Soft)"
		}
		return ""
	}()
}

func (h *Hand) calculate() (int, bool) {
	var total int
	var soft bool

	for _, card := range h.Cards {
		if !card.FaceUp {
			continue
		}

		// Maybe have the value of the card as a property on the card
		if slices.Contains([]c.Value{c.Jack, c.Queen, c.King}, card.Value) { // Face card
			total += 10
		} else if card.Value == c.Ace { // Ace
			if total+11 > 21 {
				total++
			} else {
				total += 11
				soft = true
			}
		} else { // Numbered values
			switch card.Value {
			case c.Two:
				total += 2
			case c.Three:
				total += 3
			case c.Four:
				total += 4
			case c.Five:
				total += 5
			case c.Six:
				total += 6
			case c.Seven:
				total += 7
			case c.Eight:
				total += 8
			case c.Nine:
				total += 9
			case c.Ten:
				total += 10
			}
		}

		if total > 21 && soft {
			total -= 10
			soft = false
		}
	}
	return total, soft
}

func (h *Hand) Value() int {
	value, _ := h.calculate()
	return value
}

func (h *Hand) IsSoft() bool {
	_, soft := h.calculate()
	return soft
}

func (h *Hand) IsBusted() bool {
	return h.Value() > 21
}

func (h Hand) IsBlackJack() bool {
	if len(h.Cards) == 2 && h.Value() == 21 {
		return true
	}
	return false
}

// TODO: dont know if this should be here
// Hit adds a card to the hand
func (h *Hand) Hit(c c.Card) {
	if !h.Locked {
		h.Cards = append(h.Cards, c)
	}
}

func (h *Hand) Stand() {
	h.Locked = true
}
