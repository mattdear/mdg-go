package main

import "math/rand"

type Dice struct {
    value  int
    isHold bool
}

func (dice *Dice) Roll() {
    dice.value = rand.Intn(7)
}