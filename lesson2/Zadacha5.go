package main

import "fmt"

func zadacha5() {

	fmt.Println("Дано длина ребра куба")
	var a float64
	fmt.Scan(&a)
	fmt.Println("Найти  объем куба и площадь поверхности")

	V := (a * a * a)
	S := (6 * a * a)

	fmt.Println(V)
	fmt.Println(S)
}
