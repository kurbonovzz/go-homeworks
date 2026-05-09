package main

import "fmt"

func zadacha10() {

	fmt.Println("Введите два ненулевых числа ")

	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println("Найдите сумму, разность, произведение и частное их квадратов.")

	S := (a * a) + (a * a)
	R := (a * a) - (b * b)
	P := (a * a) * (b * b)
	Ch := (a * a) / (b * b)

	fmt.Println(S, R, P, Ch)
}
