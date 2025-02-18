package player

// Player represents a player in the game
type Player struct {
	Name    string
	Hands   []Hand
	Credits float32
}

func NewPlayer(name string) Player {
	return Player{
		Name:    name,
		Credits: 100,
		Hands:   []Hand{},
	}
}

func NewDealer() Player {
	return Player{
		Name:  "Dealer",
		Hands: []Hand{},
	}
}

func (p Player) String() string {
	return p.Name
}

func (p *Player) ResetHands() {
	p.Hands = []Hand{}
}

func (p *Player) PlaceBet(bet float32) {
	hand := Hand{
		Bet: bet,
	}
	p.Hands = append(p.Hands, hand)
}
