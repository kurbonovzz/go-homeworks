package main

import "fmt"

func zadacha4() {

	fmt.Println("Дан диаметр:")
	var d, L float64
	fmt.Scan(&d)

	fmt.Println("Найти ее длину ")

	P := (3.14)

	L = P * d
	fmt.Println(L)

}
