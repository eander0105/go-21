package game

import (
	"fmt"
	"strings"

	. "github.com/eander0105/go-21/deck"
	. "github.com/eander0105/go-21/player"
)

type Game struct {
	Deck        Deck
	Players     []Player
	Dealer      Player
	HouseProfit float32
}

func NewGame() Game {
	g := Game{}

	g.Deck = NewDeck(4, true)
	g.Dealer = NewDealer()

	return g
}

func (g *Game) Play() {
	fmt.Printf("Welcome to the g of blackjack, %s!\n", g.Players[0].Name)
	fmt.Printf("You start with %.2f credits.\n", g.Players[0].Credits)
	fmt.Println("The dealer shuffles the deck.")

	for {
		// TODO: Ask how much the player wants to bet
		g.PlayRound()
	}
}

func (g *Game) PlayRound() {

	for i := range g.Players {
		player := &g.Players[i]
		player.PlaceBet(10)
	}
	g.Dealer.PlaceBet(0) // arbitrary bet

	g.Deal()

	fmt.Printf("The dealer deals you %s, and them self %s.\n", g.Players[0].Hands[0], g.Dealer.Hands[0])

	var input string
	for _, player := range g.Players {
		for i := range player.Hands {
		handLoop:
			for {
				hand := &player.Hands[i]
				if len(player.Hands) == 1 {
					fmt.Println("Your hand:", hand)
				} else {
					fmt.Printf("Your hand %d: %s\n", i, hand)
				}
				fmt.Println("Hit(h) or stand(s)?")

				_, err := fmt.Scanln(&input) // User input

				if err != nil { // Invalid input
					fmt.Println("Invalid input. Please try again.")
					continue
				}

				switch strings.ToLower(input) {
				case "h", "hit":
					card := g.Deck.DrawTopCard(true)
					hand.Hit(card)
				case "s", "stand":
					fmt.Println("You stand.")
					break handLoop
				default:
					fmt.Println("Invalid input. Please try again.")
				}
			}
		}
	}
	g.EndRound()
}

func (g *Game) EndRound() {
	dealerTotal := g.Dealer.Hands[0].Value()
	for i := range g.Players {
		player := &g.Players[i]
		for j := range player.Hands {
			hand := &player.Hands[j]
			handTotal := hand.Value()

			var payout float32

			if handTotal == dealerTotal { // Hand is a draw
				payout = hand.Bet * 1
			} else if hand.IsBlackJack() { // Hand is blackjack
				payout = hand.Bet * 2.5
			} else if handTotal <= 21 { // Hand is legal
				if handTotal > dealerTotal { // Player wins
					payout = hand.Bet * 2
				} else { // Dealer wins
					payout = 0
				}
			}

			player.Credits += payout
			g.HouseProfit -= payout
		}

		player.ResetHands()
	}
	g.Dealer.ResetHands()
}

func (g *Game) AddPlayer(name string) {
	player := NewPlayer(name)
	g.Players = append(g.Players, player)
}

func (g *Game) Deal() {
	deckIndex := 0

	// Give each player two cards
	for i := 0; i < 2; i++ {
		// Players cards
		for j := range g.Players {
			player := &g.Players[j]

			card := g.Deck[deckIndex]
			card.FaceUp = true

			player.Hands[0].AddCard(card)
			deckIndex++
		}

		card := g.Deck[deckIndex]
		if i == 0 {
			card.FaceUp = true
		} else {
			card.FaceUp = true
		}
		g.Dealer.Hands[0].AddCard(card)
		deckIndex++
	}
	g.Deck = g.Deck[deckIndex:]
}
