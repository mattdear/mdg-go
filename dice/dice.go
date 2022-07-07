package dice

import "math/rand"

type Dice struct {
	value  int
	isHold bool
}

func Roll() int {
	return rand.Intn(7)
}
