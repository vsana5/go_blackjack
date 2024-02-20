package main

import (
	"fmt"
)

func main() {

	deck := getDefaultDeck()
	deck.Shuffle()

	yourHand := Hand {true, []string{}}
	dealerHand := Hand {false, []string{}}

	yourHand.Add(deck.Draw())
	dealerHand.Add(deck.Draw())
	yourHand.Add(deck.Draw())
	dealerHand.Add(deck.Draw())

	fmt.Printf("Dealer's hand: %v\n", dealerHand)

	if yourHand.Evaluate() == 21 {
		// Autowin
		fmt.Printf("You %v got 21 you have won 🎉\n", yourHand)
		return
	}
	
	userInput:
	for {
		var command string
		fmt.Printf("Your hand %v, Hit or stay ?\n", yourHand)
		fmt.Scan(&command)

		switch command {
		case "Hit", "hit":
			yourHand.Add(deck.Draw())
			if 21 <= yourHand.Evaluate() {
				break userInput
			} else {
				fmt.Printf("Your hand: %v\n", yourHand)
			}
			break;
		case "Stay", "stay":
			break userInput
		default:
			fmt.Print("Please enter one of the commands: ")
			break
		}
	}

	if 21 < yourHand.Evaluate() {
		fmt.Printf("You busted %v and lost\n", yourHand)
		return
	}

	if yourHand.Evaluate() == 21 {
		fmt.Printf("You %v got 21 you have won 🎉\n", yourHand)
		return
	}


	for dealerHand.Evaluate() < 17 {
		dealerHand.Add(deck.Draw())
	}

	dealerHand.isOpen = true

	if 21 < dealerHand.Evaluate() {
		fmt.Printf("Dealer busted %v, You won!\n", dealerHand)
	} else if dealerHand.Evaluate() < yourHand.Evaluate() {
		fmt.Printf("You won your hand: %v vs dealer's hand %v\n", yourHand, dealerHand)
	} else {
		fmt.Printf("You Lost, your hand: %v vs dealer's hand %v\n", yourHand, dealerHand)
	}
}