package main

import (
	"fmt"
	"math/rand"
)

type Hero struct {
	Name   string
	Health int
}

func (h *Hero) Attack() {
	damage := rand.Intn(100)
	h.Health -= damage
}

func (h *Hero) Heal() {
	hp := rand.Intn(100)
	h.Health += hp

	if h.Health > 100 {
		h.Health = 100
	}
}

func main() {
	hero := Hero{
		Name:   "Prototype",
		Health: 100,
	}

	fmt.Println("Start:", hero)

	hero.Attack()
	fmt.Println("After fight:", hero)

	hero.Heal()
	fmt.Println("After heal:", hero)
}
