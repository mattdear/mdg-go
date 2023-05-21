package main

import "math/rand"

type Dice struct {
	value  int
	isHold bool
}

func (dice *Dice) Roll() {
	if dice.isHold {
		return
	}
	dice.value = rand.Intn(6) + 1
}
