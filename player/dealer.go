package player

type Dealer struct {
	Player
	hitUntil int
}

func NewDealer() Dealer {
	return Dealer{
		Player: Player{
			Name:  "Dealer",
			Hands: []Hand{},
		},
	}
}

// Wether the dealer should hit or stand
func (d *Dealer) HitOnHand(h Hand) bool {
	return h.Value() < d.hitUntil
}
