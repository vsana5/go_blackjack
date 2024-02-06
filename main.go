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

var cardToValue = map[string]int{
    "A": 11,
	"K": 10,
	"Q": 10,
	"J": 10,
	"10": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

func main() {
	shuffledDeck := shuffle(deck)
	var card string
	
	var dealerHand []string
	var yourHand []string

	card, shuffledDeck = draw(shuffledDeck, 0)
	yourHand = append(yourHand, card)
	
	card, shuffledDeck = draw(shuffledDeck, 0)
	dealerHand = append(dealerHand, card)
	
	card, shuffledDeck = draw(shuffledDeck, 0)
	yourHand = append(yourHand, card)

	card, shuffledDeck = draw(shuffledDeck, 0)
	dealerHand = append(dealerHand, card)

	dealerHandToDisplay := make([]string, len(dealerHand))
	
	for cardIdx := range dealerHandToDisplay {
        dealerHandToDisplay[cardIdx] = "?"
    }

	dealerHandToDisplay[0] = dealerHand[0]

	fmt.Printf( "Dealer's hand:\t%s\tevaluation: %d\n", dealerHandToDisplay, evaluateHand(dealerHand))
	fmt.Printf( "Your hand:    \t%s\tevaluation: %d\n", yourHand, evaluateHand(yourHand))
}

func draw(deck []string, cardIdx int) (string, []string) {
	switch {
	case cardIdx == 0:
		card := deck[cardIdx]
		return card, deck[1:]
	case cardIdx == len(deck) - 1:
		card := deck[cardIdx]
		return card, deck[:cardIdx]
	case cardIdx > len(deck) - 1 || cardIdx < 0:  
		panic(fmt.Sprintf("cardIdx [%d] is out of range for deck %s", cardIdx, deck))
	}
	card := deck[cardIdx]
	return card, append(deck[0:cardIdx], deck[cardIdx+1:]...) 
}



func evaluateHand(hand []string) int {
	result := 0
	for cardIdx := range hand {
		result += cardToValue[hand[cardIdx]]
	}
	return result
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
