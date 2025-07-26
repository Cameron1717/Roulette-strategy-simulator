package main

import "math/rand"

type wheel struct {
	Place [37]string
}

func (Roulette *wheel) IntitalizeWheel() {
	Roulette.Place[0] = "green"
	for i := 1; i < 11; i++ {
		if i%2 == 0 {
			Roulette.Place[i] = "black"
		} else {
			Roulette.Place[i] = "red"
		}
	}
	for i := 11; i < 19; i++ {
		if i%2 == 0 {
			Roulette.Place[i] = "red"
		} else {
			Roulette.Place[i] = "black"
		}
	}
	for i := 19; i < 29; i++ {
		if i%2 == 0 {
			Roulette.Place[i] = "black"
		} else {
			Roulette.Place[i] = "red"
		}
	}
	for i := 29; i < 37; i++ {
		if i%2 == 0 {
			Roulette.Place[i] = "red"
		} else {
			Roulette.Place[i] = "black"
		}
	}

}

func (Roulette *wheel) Spin() (int, string) {
	position := rand.Int() % 37
	return position, Roulette.Place[position]
}
