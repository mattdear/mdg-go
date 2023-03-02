package main

import "strconv"

func main() {
    dice1 := Dice{1,false}
    dice2 := Dice{2,false}
    dice3 := Dice{3,false}
    dice4 := Dice{4,false}
    dice5 := Dice{5,false}
    dice1.Roll()
    dice2.Roll()
    dice3.Roll()
    dice4.Roll()
    dice5.Roll()


    board := Board{0,0,0,0,0,0}

    println(dice1.value)
    println(dice2.value)
    println(dice3.value)
    println(dice4.value)
    println(dice5.value)
    println(board.DisplayDice(strconv.Itoa(dice1.value),strconv.Itoa(dice2.value),strconv.Itoa(dice3.value),strconv.Itoa(dice4.value),strconv.Itoa(dice5.value)))
    println(board.diceTotal(dice1.value,dice2.value,dice3.value,dice4.value,dice5.value))

    board.activeDice(dice1.value)
    println(board.activeSixes)
}