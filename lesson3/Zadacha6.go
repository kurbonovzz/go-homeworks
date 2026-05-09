package main

import (
	"fmt"
	"math/rand"
)

type Hero struct {
	Name   string
	Health int
}

func attack(h *Hero) Hero {

	hp := rand.Intn(100)

	h.Health = h.Health - hp
	return *h

}

func heal(h *Hero) {

	hp := rand.Intn(100)

	h.Health = h.Health + hp

	if h.Health > 100 {
		h.Health = 100
	}

	return
}

func zadacha6() {

	hero := Hero{
		Name:   "Prototype",
		Health: 100,
	}
	fmt.Println("Start :", hero)

	hero = attack(&hero)
	fmt.Println(" After fight:", hero)

	heal(&hero)
	fmt.Println(" After heal:", hero)
}
