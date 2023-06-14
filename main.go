package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
)

//TODO
//Check Scores
//Need to add the ability to blank out a line on the scorecard
//Need to fix the ability to access some code logic

func main() {
	var board Board
	var scorecard Scorecard
	var dice1, dice2, dice3, dice4, dice5 Dice
	var tempDiceString string
	gameLoop := true
	var userInputLoop, rollAgainLoop, dice1HoldLoop, dice2HoldLoop, dice3HoldLoop, dice4HoldLoop, dice5HoldLoop, blankRowLoop bool
	forfeitRollAgain := false
	unusedRows := [13]int{1,1,1,1,1,1,1,1,1,1,1,1,1}
	mainMenuOptions := []int{1, 2, 3}
	scoreboardOptions := []int{1,2,3,4,5,6,7,8,9,10,11,12,13}
	yesNoOptions := []int{1,2}
	goToMainMenuOptions := []int{1}
	goToMainMenuError := "\nInput invalid, please enter 1.\n\n"
	yesNoError := "\nInput invalid, please enter 1 or 2.\n\n"
	inputInvalidError := "\nInput invalid, please try again.\n\n"
	rowError := "\nScore cannot be entered here, please choose another row.\n\n"
	blankRowError := "\nScore cannot be entered here, would you like to blank the row?.\n\n"
	oneOrTwoError := "\nPlease enter 1 or 2.\n\n"
	for {
		clearTerminal()
		println("*------- Welcome to MDG lets play! -------*\n\nPlease select one of the following options:\n1 - New Game\n2 - Exit\n")
		switch getUserInputValidated(mainMenuOptions, inputInvalidError) {
		case 1:
			gameCycles := 0
			for {
				for rollCounter := 0; rollCounter <= 2; rollCounter++ {
					if rollCounter == 2 || forfeitRollAgain {
						if !forfeitRollAgain {
							dice1.Roll()
							dice2.Roll()
							dice3.Roll()
							dice4.Roll()
							dice5.Roll()
						}
						tempDiceString = updateDiceArray(dice1.value, dice2.value, dice3.value, dice4.value, dice5.value)
						print(scorecard.Display())
						print(board.DisplayDice(strconv.Itoa(dice1.value), strconv.Itoa(dice2.value), strconv.Itoa(dice3.value), strconv.Itoa(dice4.value), strconv.Itoa(dice5.value)))
						board.activeDice(dice1.value)
						board.activeDice(dice2.value)
						board.activeDice(dice3.value)
						board.activeDice(dice4.value)
						board.activeDice(dice5.value)
						print("\nWhere would you like to place this score? (1-13)\n\n")
						userInputLoop = true
						for {
							switch getUserInputValidated(scoreboardOptions, rowError) {
							case 1:
								if scorecard.ones == 0 && board.activeOnes > 0 && unusedRows[0] == 1 {
									scorecard.ones = board.activeOnes
									userInputLoop = false
									unusedRows[0] = 0
								} else {
									print(blankRowError)
								}
							case 2:
								if scorecard.twos == 0 && board.activeTwos > 0 && unusedRows[1] == 1 {
									scorecard.twos = board.activeTwos * 2
									userInputLoop = false
									unusedRows[1] = 0
								} else {
									print(blankRowError)
								}
							case 3:
								if scorecard.threes == 0 && board.activeThrees > 0 && unusedRows[2] == 1 {
									scorecard.threes = board.activeThrees * 3
									userInputLoop = false
									unusedRows[2] = 0
								} else {
									print(blankRowError)
								}
							case 4:
								if scorecard.fours == 0 && board.activeFours > 0 && unusedRows[3] == 1 {
									scorecard.fours = board.activeFours * 4
									userInputLoop = false
									unusedRows[3] = 0
								} else {
									print(blankRowError)
								}
							case 5:
								if scorecard.fives == 0 && board.activeFives > 0 && unusedRows[4] == 1 {
									scorecard.fives = board.activeFives * 5
									userInputLoop = false
									unusedRows[4] = 0
								} else {
									print(blankRowError)
								}
							case 6:
								if scorecard.sixes == 0 && board.activeSixes > 0 && unusedRows[5] == 1 {
									scorecard.sixes = board.activeSixes * 6
									userInputLoop = false
									unusedRows[5] = 0
								} else {
									print(blankRowError)
								}
							case 7:
								if (board.activeOnes == 2 || board.activeTwos == 2 || board.activeThrees == 2 || board.activeFours == 2 || board.activeFives == 2 || board.activeSixes == 2) && (board.activeOnes == 3 || board.activeTwos == 3 || board.activeThrees == 3 || board.activeFours == 3 || board.activeFives == 3 || board.activeSixes == 3) && scorecard.twoThreeMatch == 0 && unusedRows[6] == 1 {
									scorecard.twoThreeMatch = 25
									userInputLoop = false
									unusedRows[6] = 0
								} else {
									print(blankRowError)
								}
							case 8:
								if (board.activeOnes == 3 || board.activeTwos == 3 || board.activeThrees == 3 || board.activeFours == 3 || board.activeFives == 3 || board.activeSixes == 3) && scorecard.threeMatch == 0 && unusedRows[7] == 1 {
									scorecard.threeMatch = board.diceTotal(dice1.value, dice2.value, dice3.value, dice4.value, dice5.value)
									userInputLoop = false
									unusedRows[7] = 0
								} else {
									print(blankRowError)
								}
							case 9:
								if (board.activeOnes == 4 || board.activeTwos == 4 || board.activeThrees == 4 || board.activeFours == 4 || board.activeFives == 4 || board.activeSixes == 4) && scorecard.fourMatch == 0 && unusedRows[8] == 1 {
									scorecard.fourMatch = board.diceTotal(dice1.value, dice2.value, dice3.value, dice4.value, dice5.value)
									userInputLoop = false
									unusedRows[8] = 0
								} else {
									print(blankRowError)
								}
							case 10:
								if (board.activeOnes == 5 || board.activeTwos == 5 || board.activeThrees == 5 || board.activeFours == 5 || board.activeFives == 5 || board.activeSixes == 5) && scorecard.fiveMatch == 0 && unusedRows[9] == 1 {
									scorecard.fiveMatch = 50
									userInputLoop = false
									unusedRows[9] = 0
								} else {
									print(blankRowError)
								}
							case 11:
								tempDiceStringLength := len(tempDiceString)
								if tempDiceStringLength >= 3 && scorecard.threeLine == 0 && unusedRows[10] == 1 {
									scorecard.threeLine = 30
									userInputLoop = false
									unusedRows[10] = 0
								} else {
									print(blankRowError)
								}
							case 12:
								tempDiceStringLength := len(tempDiceString)
								if tempDiceStringLength >= 4 && scorecard.fourLine == 0 && unusedRows[11] == 1 {
									scorecard.fourLine = 40
									userInputLoop = false
									unusedRows[11] = 0
								} else {
									print(blankRowError)
								}
							case 13:
								if scorecard.extras == 0 && unusedRows[12] == 1 {
									scorecard.extras = board.diceTotal(dice1.value, dice2.value, dice3.value, dice4.value, dice5.value)
									userInputLoop = false
									unusedRows[12] = 0
								} else {
									print(blankRowError)
									blankRowLoop = true
								}
							}
							if !userInputLoop {
								break
							}
						}
						scorecard.calculateTopScore()
						scorecard.calculateBottomScore()
						scorecard.calculateTotalScore()
						board.activeOnes = 0
						board.activeTwos = 0
						board.activeThrees = 0
						board.activeFours = 0
						board.activeFives = 0
						forfeitRollAgain = false
						rollCounter = 0
					}
					clearTerminal()
					dice1.Roll()
					dice2.Roll()
					dice3.Roll()
					dice4.Roll()
					dice5.Roll()
					print(scorecard.Display())
					print(board.DisplayDice(strconv.Itoa(dice1.value), strconv.Itoa(dice2.value), strconv.Itoa(dice3.value), strconv.Itoa(dice4.value), strconv.Itoa(dice5.value)))
					print("\nRound " + strconv.Itoa(gameCycles+1))
					print("\nRoll " + strconv.Itoa(rollCounter+1) + ` of 3`)
					dice1.isHold = false
					dice2.isHold = false
					dice3.isHold = false
					dice4.isHold = false
					dice5.isHold = false
					if rollCounter != 2 {
						print("\nRoll again? y(1)/n(2)\n\n")
						rollAgainLoop = true
						for {
							switch getUserInputValidated(yesNoOptions,yesNoError) {
							case 1:
								rollAgainLoop = false
							case 2:
								forfeitRollAgain = true
								rollAgainLoop = false
							if !rollAgainLoop {
								break
							}
						}
						if !forfeitRollAgain {
							print("\nHold dice 1? y(1)/n(2)\n\n")
							dice1HoldLoop = true
							for {
								switch getUserInputValidated(yesNoOptions,yesNoError) {
								case 1:
									dice1.isHold = true
									dice1HoldLoop = false
								case 2:
									dice1HoldLoop = false
								}
								if !dice1HoldLoop {
									break
								}
							}
							print("\nHold dice 2? y(1)/n(2)\n\n")
							dice2HoldLoop = true
							for {
								switch getUserInputValidated(yesNoOptions,yesNoError) {
								case 1:
									dice2.isHold = true
									dice2HoldLoop = false
								case 2:
									dice2HoldLoop = false
								}
								if !dice2HoldLoop {
									break
								}
							}
							print("\nHold dice 3? y(1)/n(2)\n\n")
							dice3HoldLoop = true
							for {
								switch getUserInputValidated(yesNoOptions,yesNoError) {
								case 1:
									dice3.isHold = true
									dice3HoldLoop = false
								case 2:
									dice3HoldLoop = false
								}
								if !dice3HoldLoop {
									break
								}
							}
							print("\nHold dice 4? y(1)/n(2)\n\n")
							dice4HoldLoop = true
							for {
								switch getUserInputValidated(yesNoOptions,yesNoError) {
								case 1:
									dice4.isHold = true
									dice4HoldLoop = false
								case 2:
									dice4HoldLoop = false
								}
								if !dice4HoldLoop {
									break
								}
							}
							print("\nHold dice 5? y(1)/n(2)\n\n")
							dice5HoldLoop = true
							for {
								switch getUserInputValidated(yesNoOptions,yesNoError) {
								case 1:
									dice5.isHold = true
									dice5HoldLoop = false
								case 2:
									dice5HoldLoop = false
								}
								if !dice5HoldLoop {
									break
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
						print("\nGo back to the main menumain menu y(1).\n\n")
						if getUserInputValidated(goToMainMenuOptions, goToMainMenuError) != 0 {
							break
						}
					}
					gameCycles++
				}
			}
		case 2:
			println("\nThanks for playing MDG, have a great day. :)\n\n")
			gameLoop = false
		}
		if !gameLoop {
			break
		}
	}
}

func getUserInputValidated(userOptions[]int, errorMessage string) int {
	var userInput int
	userInputLoop := true
	for {
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			return 0
		}
		for _, option := range userOptions{
			if option == userInput {
				userInputLoop = false
			}
		}
		if !userInputLoop {
			break
		}
		println(errorMessage)
	}
	return userInput
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return
	}
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
