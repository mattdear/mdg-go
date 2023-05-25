package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	var userInput int
	var board Board
	var scorecard Scorecard
	var dice1, dice2, dice3, dice4, dice5 Dice
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
					forfeitRollAgain := false
					dice1.isHold = false
					dice2.isHold = false
					dice3.isHold = false
					dice4.isHold = false
					dice5.isHold = false
					if rollCounter == 2 || forfeitRollAgain {
						print("\nWhere would you like to place this score? (1-13)\n\n")
						switch getUserInput() {
						case 1:
							scorecard.ones = 1
						case 2:
							scorecard.twos = 2
						case 3:
							scorecard.threes = 3
						case 4:
							scorecard.fours = 4
						case 5:
							scorecard.fives = 5
						case 6:
							scorecard.sixes = 6
						case 7:
							scorecard.twoThreeMatch = 23
						case 8:
							scorecard.threeMatch = 33
						case 9:
							scorecard.fourMatch = 44
						case 10:
							scorecard.fiveMatch = 55
						case 11:
							scorecard.threeLine = 333
						case 12:
							scorecard.fourLine = 4444
						case 13:
							scorecard.extras = 13
						}
						break
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
					break
				}
				gameCycles++
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

// //Setting active dice
// board.activeDice(board, d1.getValue());
// board.activeDice(board, d2.getValue());
// board.activeDice(board, d3.getValue());
// board.activeDice(board, d4.getValue());
// board.activeDice(board, d5.getValue());
// System.out.println(scorecard.display());
// System.out.println(board.diceValues(d1.getValue(), d2.getValue(), d3.getValue(), d4.getValue(), d5.getValue()));
// System.out.println("Round " + gameCycles + " of 13");
// boolean isScorecardLocation = true;

// //Scorecard location loop
// do {
// System.out.println("Place on scorecard? (1-13)");
// while (!numberScanner.hasNextInt()) {
// System.out.println("\nInput invalid, please try again.\n\n");
// System.out.println("Place on scorecard? (1-13)");
// numberScanner.next();
// }
// userInputNumber = numberScanner.nextInt();
// int diceTotal = board.diceTotal(d1.getValue(), d2.getValue(), d3.getValue(), d4.getValue(), d5.getValue());

// //Dice added to array to be sorted
// int[] diceArray = new int[5];
// diceArray[0] = d1.getValue();
// diceArray[1] = d2.getValue();
// diceArray[2] = d3.getValue();
// diceArray[3] = d4.getValue();
// diceArray[4] = d5.getValue();
// Arrays.sort(diceArray);

// //Looks at first array item and compares to the second if they are not the same the first array item will be added to the string.
// //This continues though the array until the penultimate member. The last array member is added by default after the loop has ended.
// String tempDiceString = "";
// for (int i = 0; i<4; i++){
// if(diceArray[i] != diceArray[i+1]){
// tempDiceString = tempDiceString + diceArray[i];
// }
// }
// tempDiceString = tempDiceString + diceArray[4];

// //Scorecard input verification
// if (userInputNumber == 1) { //Ones
// if (scorecard.getOnes() == 0 && board.getActiveOnes() > 0) {
// scorecard.setOnes(board.getActiveOnes()*1);
// } else {
// scorecard.setOnes(0);
// System.out.println("\nThis scorecard row has been zeroed due to invalid dice.\n\n");
// }
// isScorecardLocation = false;
// } else if (userInputNumber == 2) { //Twos
// if (scorecard.getTwos() == 0 && board.getActiveTwos() > 0) {
// scorecard.setTwos(board.getActiveTwos()*2);
// } else {
// scorecard.setTwos(0);
// System.out.println("\nThis scorecard row has been zeroed due to invalid dice.\n\n");
// }
// isScorecardLocation = false;
// } else if (userInputNumber == 3) { //Threes
// if (scorecard.getThrees() == 0 && board.getActiveThrees() > 0) {
// scorecard.setThrees(board.getActiveThrees()*3);
// } else {
// scorecard.setThrees(0);
// System.out.println("\nThis scorecard row has been zeroed due to invalid dice.\n\n");
// }
// isScorecardLocation = false;
// } else if (userInputNumber == 4) { //Fours
// if (scorecard.getFours() == 0 && board.getActiveFours() > 0) {
// scorecard.setFours(board.getActiveFours()*4);
// } else {
// scorecard.setFours(0);
// System.out.println("\nThis scorecard row has been zeroed due to invalid dice.\n\n");
// }
// isScorecardLocation = false;
// } else if (userInputNumber == 5) { //Fives
// if (scorecard.getFives() == 0 && board.getActiveFives() > 0) {
// scorecard.setFives(board.getActiveFives()*5);
// } else {
// scorecard.setFives(0);
// System.out.println("\nThis scorecard row has been zeroed due to invalid dice.\n\n");
// }
// isScorecardLocation = false;
// } else if (userInputNumber == 6) { //Sixes
// if (scorecard.getSixes() == 0 && board.getActiveSixes() > 0) {
// scorecard.setSixes(board.getActiveSixes()*6);
// } else {
// scorecard.setSixes(0);
// System.out.println("\nThis scorecard row has been zeroed due to invalid dice.\n\n");
// }
// isScorecardLocation = false;
// } else if (userInputNumber == 7) { //2 3 Match
// if ((board.getActiveOnes() == 2 || board.getActiveTwos() == 2 || board.getActiveThrees() == 2 || board.getActiveFours() == 2 || board.getActiveFives() == 2 || board.getActiveSixes() == 2) && (board.getActiveOnes() == 3 || board.getActiveTwos() == 3 || board.getActiveThrees() == 3 || board.getActiveFours() == 3 || board.getActiveFives() == 3 || board.getActiveSixes() == 3) && scorecard.getTwoThreeMatch() == 0) {
// scorecard.setTwoThreeMatch(25);
// } else {
// scorecard.setTwoThreeMatch(0);
// System.out.println("\nThis scorecard row has been zeroed due to invalid dice.\n\n");
// }
// isScorecardLocation = false;
// } else if (userInputNumber == 8) { //3 Match
// if ((board.getActiveOnes() == 3 || board.getActiveTwos() == 3 || board.getActiveThrees() == 3 || board.getActiveFours() == 3 || board.getActiveFives() == 3 || board.getActiveSixes() == 3) && scorecard.getThreeMatch() == 0) {
// scorecard.setThreeMatch(diceTotal);
// } else {
// scorecard.setThreeMatch(0);
// System.out.println("\nThis scorecard row has been zeroed due to invalid dice.\n\n");
// }
// isScorecardLocation = false;
// } else if (userInputNumber == 9) { //4 Match
// if ((board.getActiveOnes() == 4 || board.getActiveTwos() == 4 || board.getActiveThrees() == 4 || board.getActiveFours() == 4 || board.getActiveFives() == 4 || board.getActiveSixes() == 4) && scorecard.getFourMatch() == 0) {
// scorecard.setFourMatch(diceTotal);
// } else {
// scorecard.setFourMatch(0);
// System.out.println("\nThis scorecard row has been zeroed due to invalid dice.\n\n");
// }
// isScorecardLocation = false;
// } else if (userInputNumber == 10) { //5 Match
// if ((board.getActiveOnes() == 5 || board.getActiveTwos() == 5 || board.getActiveThrees() == 5 || board.getActiveFours() == 5 || board.getActiveFives() == 5 || board.getActiveSixes() == 5) && scorecard.getFiveMatch() == 0) {
// scorecard.setFiveMatch(50);
// } else {
// scorecard.setFiveMatch(0);
// System.out.println("\nThis scorecard row has been zeroed due to invalid dice.\n\n");
// }
// isScorecardLocation = false;
// } else if (userInputNumber == 11) { //3 Line
// if (tempDiceString.length() >= 3 && scorecard.getThreeLine() == 0) {
// scorecard.setThreeLine(30);
// } else {
// scorecard.setThreeLine(0);
// System.out.println("\nThis scorecard row has been zeroed due to invalid dice.\n\n");
// }
// isScorecardLocation = false;
// } else if (userInputNumber == 12) { //4 Line
// if (tempDiceString.length() >= 4 && scorecard.getFourLine() == 0) {
// scorecard.setFourLine(40);
// } else {
// scorecard.setFourLine(0);
// System.out.println("\nThis scorecard row has been zeroed due to invalid dice.\n\n");
// }
// isScorecardLocation = false;
// } else if (userInputNumber == 13) { //Extras
// if (scorecard.getExtras() == 0) {
// scorecard.setExtras(diceTotal);
// } else {
// scorecard.setExtras(0);
// System.out.println("\nThis scorecard row has been zeroed due to invalid dice.\n\n");
// }
// isScorecardLocation = false;
// } else {
// System.out.println("\nInput invalid, please try again.\n\n");
// }

// //Variable reset
// diceTotal = 0;
// } while (isScorecardLocation == true);
// scorecard.calcTopScore();
// scorecard.calcBottomScore();
// scorecard.calcTotalScore();

// //Veriable resets
// d1.setHold(false);
// d2.setHold(false);
// d3.setHold(false);
// d4.setHold(false);
// d5.setHold(false);
// board.setActiveOnes(0);
// board.setActiveTwos(0);
// board.setActiveThrees(0);
// board.setActiveFours(0);
// board.setActiveFives(0);
// board.setActiveSixes(0);
// gameCycles = gameCycles + 1;
// } while (gameCycles < 14);
// System.out.println(scorecard.display());
// System.out.println("Your final score " + scorecard.getTotalScore() + ", Thanks for playing.\n\n\n");

// 	println(dice1.value)
// 	println(dice2.value)
// 	println(dice3.value)
// 	println(dice4.value)
// 	println(dice5.value)
// 	println(board.DisplayDice(strconv.Itoa(dice1.value), strconv.Itoa(dice2.value), strconv.Itoa(dice3.value), strconv.Itoa(dice4.value), strconv.Itoa(dice5.value)))
// 	println(board.diceTotal(dice1.value, dice2.value, dice3.value, dice4.value, dice5.value))

// 	board.activeDice(dice1.value)
// 	println(board.activeSixes)
// 	println(scorecard.DisplayScorecard())
// }
