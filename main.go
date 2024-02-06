package main

import (
    "fmt"
    "math/rand"
)

var deck = []string{
	"A", "A", "A", "A", 
	"K", "K", "K", "K", 
	"Q", "Q", "Q", "Q",
	"J", "J", "J", "J",
	"10", "10", "10", "10",
	"9", "9", "9", "9",
	"8", "8", "8", "8",
	"7", "7", "7", "7",
	"6", "6", "6", "6",
	"5", "5", "5", "5",
	"4", "4", "4", "4",
	"3", "3", "3", "3",
	"2", "2", "2", "2",
}

var dealerHand []string
var yourHand []string

func main() {
	shuffledDeck := shuffle(deck)
	var card string

	card, shuffledDeck = draw(shuffledDeck, 0)
	dealerHand = append(dealerHand, card)
	card, shuffledDeck = draw(shuffledDeck, 0)
	yourHand = append(yourHand, card)
	card, shuffledDeck = draw(shuffledDeck, 0)
	dealerHand = append(dealerHand, card)
	card, shuffledDeck = draw(shuffledDeck, 0)
	yourHand = append(yourHand, card)
	fmt.Printf( "Dealer's hand: %s\n Your hand: %s\n", dealerHand, yourHand)
}

func draw(deck []string, card_idx int) (string, []string) {
	switch {
	case card_idx == 0:
		card := deck[card_idx]
		return card, deck[1:]
	case card_idx == len(deck) - 1:
		card := deck[card_idx]
		return card, deck[:card_idx]
	case card_idx > len(deck) - 1 || card_idx < 0:  
		panic(fmt.Sprintf("card_idx [%d] is out of range for deck %s", card_idx, deck))
	}
	card := deck[card_idx]
	return card, append(deck[0:card_idx], deck[card_idx+1:]...) 
}

func shuffle(deck []string) []string {
	shuffledDeck := make([]string, len(deck))

	copy(shuffledDeck, deck)

	for i := range shuffledDeck {
		j := rand.Intn(i + 1)
		shuffledDeck[i], shuffledDeck[j] = shuffledDeck[j], shuffledDeck[i]
	}

	return shuffledDeck
}

