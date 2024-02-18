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
	"A":  11,
	"K":  10,
	"Q":  10,
	"J":  10,
	"10": 10,
	"9":  9,
	"8":  8,
	"7":  7,
	"6":  6,
	"5":  5,
	"4":  4,
	"3":  3,
	"2":  2,
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

	fmt.Printf("Dealer's hand:\t%s\tevaluation: %d\n", dealerHandToDisplay, evaluateHand(dealerHand))
	fmt.Printf("Your hand:    \t%s\tevaluation: %d\n", yourHand, evaluateHand(yourHand))

	yourScore := evaluateHand(yourHand)
	
	if yourScore == 21 {
		// Autowin
		fmt.Println("You got 21 you have won 🎉")
		return
	}
		
	userGetCards(yourHand, shuffledDeck)

	if(21 < evaluateHand(yourHand)) {
		fmt.Println("You've busted and lost")
		return
	}

	for evaluateHand(dealerHand) < 17 {
		card, shuffledDeck = draw(shuffledDeck, 0)
		dealerHand = append(dealerHand, card)
	}

	if 21 < evaluateHand(dealerHand) {
		fmt.Println("Dealer busted, You won!")
	} else if evaluateHand(dealerHand) < evaluateHand(yourHand) {
		fmt.Println("You won")
	} else {
		fmt.Printf("You Lost, dealer's hand: %s with evaluation: %d \n", dealerHand, evaluateHand(dealerHand))
	}
}

func userGetCards(yourHand, shuffledDeck []string) {
	var card string
	for {
		var command string
		fmt.Println("hit or stay")
		fmt.Scan(&command)

		switch command {
		case "Hit", "hit":
			card, shuffledDeck = draw(shuffledDeck, 0)
			yourHand = append(yourHand, card)
			if(21 < evaluateHand(yourHand)) {
				return
			} else {
				fmt.Printf("Your hand:    \t%s\tevaluation: %d\n", yourHand, evaluateHand(yourHand))
			}
			break;
		case "Stay", "stay":
			return;
		default:
			fmt.Print("Please enter one of the commands: ")
			break
		}
	}
}

func draw(deck []string, cardIdx int) (string, []string) {
	switch {
	case cardIdx == 0:
		card := deck[cardIdx]
		return card, deck[1:]
	case cardIdx == len(deck)-1:
		card := deck[cardIdx]
		return card, deck[:cardIdx]
	case cardIdx > len(deck)-1 || cardIdx < 0:
		panic(fmt.Sprintf("cardIdx [%d] is out of range for deck %s", cardIdx, deck))
	}
	card := deck[cardIdx]
	return card, append(deck[0:cardIdx], deck[cardIdx+1:]...)
}

func evaluateHand(hand []string) int {
	result := 0

	acesCount := 0

	for cardIdx := range hand {
		card := hand[cardIdx]
		if card == "A" {
			acesCount++
		}
		result += cardToValue[card]
	}

	if result > 21 && acesCount > 0 {
		result -= min((result-12)/10, acesCount) * 10
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
