package main

type Board struct {
    activeOnes int
    activeTwos int
    activeThrees int
    activeFours int
    activeFives int
    activeSixes int
}


func (board *Board) DisplayDice(diceOne string, diceTwo string, diceThree string, diceFour string, diceFive string) string {
    output := "\n| Roll --------------------------|\n| Dice  | 1  | 2  | 3  | 4  | 5  |\n" +
                "| Value | " + diceOne + "  | " + diceTwo + "  | " + diceThree + "  | " + diceFour + "  | " + diceFive + "  |\n"
    return output
}

func (board *Board) diceTotal(diceOne int, diceTwo int, diceThree int, diceFour int, diceFive int) int {
    diceTotal := diceOne + diceTwo + diceThree + diceFour + diceFive
    return diceTotal
}

func (board *Board) activeDice(value int) {
    switch {
        case value == 1:
            board.activeOnes = board.activeOnes + 1
            break
        case value == 2:
            board.activeTwos = board.activeOnes + 1
            break
        case value == 3:
            board.activeThrees = board.activeThrees + 1
            break
        case value == 4:
            board.activeFours = board.activeFours + 1
            break
        case value == 5:
            board.activeFives = board.activeFives + 1
            break
    }
}