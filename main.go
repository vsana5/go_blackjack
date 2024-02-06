package main

import "fmt"

var deck = []string{
	"A", 
	"K", 
	"Q",
	"J",
	"10",
	"9",
	"8",
	"7",
	"6",
	"5",
	"4",
	"3",
	"2",
}

func main() {
	for len(deck) > 0 {
		var card string
		card, deck = draw(deck, len(deck) - 1)
		fmt.Printf( "%s, %s\n", card, deck)
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