package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
)

func main() {
	var board Board
	var scorecard Scorecard
	var dice1, dice2, dice3, dice4, dice5 Dice
	var tempDiceString string
	gameLoop := true
	for {
		clearTerminal()
		println("*------- Welcome to MDG lets play! -------*\n\nPlease select one of the following options:\n1 - New Game\n2 - Exit\n")
		switch getUserInput() {
		case 1:
			gameCycles := 0
			for {
				for rollCounter := 0; rollCounter <= 2; rollCounter++ {
					clearTerminal()
					dice1.Roll()
					dice2.Roll()
					dice3.Roll()
					dice4.Roll()
					dice5.Roll()
					print(scorecard.Display())
					print(board.DisplayDice(strconv.Itoa(dice1.value), strconv.Itoa(dice2.value), strconv.Itoa(dice3.value), strconv.Itoa(dice4.value), strconv.Itoa(dice5.value)))
					print("\nRound " + strconv.Itoa(gameCycles+1))
					print("\nRoll " + strconv.Itoa(rollCounter+1) + " of 3")
//					forfeitRollAgain := false
					dice1.isHold = false
					dice2.isHold = false
					dice3.isHold = false
					dice4.isHold = false
					dice5.isHold = false
					if rollCounter == 2 {
						tempDiceString = updateDiceArray(dice1.value, dice2.value, dice3.value, dice4.value, dice5.value)
						print("\nWhere would you like to place this score? (1-13)\n\n")
						switch getUserInput() {
						case 1:
							if scorecard.ones == 0 && board.activeOnes > 0 {
								scorecard.ones = board.activeOnes
							} else {
								scorecard.ones = 0
							}
						case 2:
							if scorecard.twos == 0 && board.activeTwos > 0 {
								scorecard.twos = board.activeTwos * 2
							} else {
								scorecard.twos = 0
							}
						case 3:
							if scorecard.threes == 0 && board.activeThrees > 0 {
								scorecard.threes = board.activeThrees * 3
							} else {
								scorecard.threes = 0
							}
						case 4:
							if scorecard.fours == 0 && board.activeFours > 0 {
								scorecard.fours = board.activeFours * 4
							} else {
								scorecard.fours = 0
							}
						case 5:
							if scorecard.fives == 0 && board.activeFives > 0 {
								scorecard.fives = board.activeFives * 5
							} else {
								scorecard.fives = 0
							}
						case 6:
							if scorecard.sixes == 0 && board.activeSixes > 0 {
								scorecard.sixes = board.activeSixes * 6
							} else {
								scorecard.sixes = 0
							}
						case 7:
							if (board.activeOnes == 2 || board.activeTwos == 2 || board.activeThrees == 2 || board.activeFours == 2 || board.activeFives == 2 || board.activeSixes == 2) && (board.activeOnes == 3 || board.activeTwos == 3 || board.activeThrees == 3 || board.activeFours == 3 || board.activeFives == 3 || board.activeSixes == 3) && scorecard.twoThreeMatch == 0 {
								scorecard.twoThreeMatch = 25
							} else {
								scorecard.twoThreeMatch = 0
							}
						case 8:
							if (board.activeOnes == 3 || board.activeTwos == 3 || board.activeThrees == 3 || board.activeFours == 3 || board.activeFives == 3 || board.activeSixes == 3) && scorecard.threeMatch == 0 {
								scorecard.threeMatch = board.diceTotal(dice1.value, dice2.value, dice3.value, dice4.value, dice5.value)
							} else {
								scorecard.threeMatch = 0
							}
						case 9:
							if (board.activeOnes == 4 || board.activeTwos == 4 || board.activeThrees == 4 || board.activeFours == 4 || board.activeFives == 4 || board.activeSixes == 4) && scorecard.fourMatch == 0 {
								scorecard.fourMatch = board.diceTotal(dice1.value, dice2.value, dice3.value, dice4.value, dice5.value)
							} else {
								scorecard.fourMatch = 0
							}
						case 10:
							if (board.activeOnes == 5 || board.activeTwos == 5 || board.activeThrees == 5 || board.activeFours == 5 || board.activeFives == 5 || board.activeSixes == 5) && scorecard.fiveMatch == 0 {
								scorecard.fiveMatch = 50
							} else {
								scorecard.fiveMatch = 0
							}
						case 11:
							tempDiceStringLength := len(tempDiceString)
							if tempDiceStringLength >= 3 && scorecard.threeLine == 0 {
								scorecard.threeLine = 30
							} else {
								scorecard.threeLine = 0
							}
						case 12:
							tempDiceStringLength := len(tempDiceString)
							if tempDiceStringLength >= 4 && scorecard.fourLine == 0 {
								scorecard.fourLine = 40
							} else {
								scorecard.fourLine = 0
							}
						case 13:
							if scorecard.extras == 0 {
								scorecard.extras = board.diceTotal(dice1.value, dice2.value, dice3.value, dice4.value, dice5.value)
							} else {
								scorecard.extras = 0
							}
						}
						if rollCounter != 2 {
							print("\nRoll again? y(1)/n(2)\n\n")
							if getUserInput() == 2 {
								forfeitRollAgain = true
							}
							if !forfeitRollAgain {
								print("\nHold dice 1? y(1)/n(2)\n\n")
								if getUserInput() == 1 {
									dice1.isHold = true
								}
								print("\nHold dice 2? y(1)/n(2)\n\n")
								if getUserInput() == 1 {
									dice2.isHold = true
								}
								print("\nHold dice 3? y(1)/n(2)\n\n")
								if getUserInput() == 1 {
									dice3.isHold = true
								}
								print("\nHold dice 4? y(1)/n(2)\n\n")
								if getUserInput() == 1 {
									dice4.isHold = true
								}
								print("\nHold dice 5? y(1)/n(2)\n\n")
								if getUserInput() == 1 {
									dice5.isHold = true
								}
							}
						}
					}
					dice1.isHold = false
					dice2.isHold = false
					dice3.isHold = false
					dice4.isHold = false
					dice5.isHold = false
					scorecard.calculateTopScore()
					scorecard.calculateBottomScore()
					scorecard.calculateTotalScore()
					if gameCycles == 12 {
						clearTerminal()
						print(scorecard.Display())
						print("\nGAME OVER")
						print("\nPress enter to go to the main menu.\n\n")
						if getUserInput() != 0 {
							break
						}
					}
					gameCycles++
				}
			}
		case 2:
			println("\nThanks for playing MDG, have a great day. :)\n\n")
			gameLoop = false
		default:
			println("\nInput invalid, please try again.\n\n")
		}
		if !gameLoop {
			break
		}
	}
}

func getUserInput() int {
	var userInput int
	fmt.Scanln(&userInput)
	return userInput
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func updateDiceArray(dice1, dice2, dice3, dice4, dice5 int) string {
	diceValueArray := []int{dice1, dice2, dice3, dice4, dice5}
	sort.Ints(diceValueArray)
	return createDiceString(diceValueArray)
}

func createDiceString(diceValueArray []int) string {
	tempDiceString := ""
	for diceCounter := 0; diceCounter < 4; diceCounter++ {
		if diceValueArray[diceCounter] != diceValueArray[diceCounter+1] {
			tempDiceString = tempDiceString + strconv.Itoa(diceValueArray[diceCounter])
		}
		tempDiceString = tempDiceString + strconv.Itoa(diceValueArray[4])
	}
	return tempDiceString
}
