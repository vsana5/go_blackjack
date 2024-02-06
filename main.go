package main

import "fmt"

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

func main() {
	gameDeck := make([]string, 52)

	copy(gameDeck, deck)

	for len(gameDeck) > 0 {
		var card string
		card, gameDeck = draw(gameDeck, len(gameDeck) - 1)
		fmt.Printf( "%s, %s\n", card, gameDeck)
	}
	fmt.Println("The End")
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