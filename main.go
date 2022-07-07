package main

import (
	"fmt"

	"github.com/mattdear/mdg-go/dice"
)

func main() {
	var stopGame bool
	var mainMenuInput int

	for {
		if stopGame {
			break
		}
		fmt.Println("\n--- Welcome to MDG lets play! ---\n\nPlease select one of the following options:\n1 - New Game\n2 - Exit")
		fmt.Scanln(&mainMenuInput)
		switch mainMenuInput {
		case 1: // New game
			value := dice.Roll()
			fmt.Println(value)
		case 2: // Exit
			fmt.Println("\nThanks for playing MDG, have a great day.")
			stopGame = true
		default:
			fmt.Println("\nInput invalid, please try again.")
		}
	}
}
