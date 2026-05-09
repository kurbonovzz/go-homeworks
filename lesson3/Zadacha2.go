package main

import "fmt"

func zadacha7() {

	fmt.Println("Введите длину l:")
	var L float64
	fmt.Scanln(&L)

	fmt.Println("Введите площадь круга S:")
	var S int
	fmt.Scanln(&S)

	fmt.Println("Введите радиуса R:")
	var R float64
	fmt.Scanln(&R)

	P := 3.14

	fmt.Println("Ответ длина окружности:")
	L = 2 * P * R
	fmt.Println(L)

	fmt.Println("Ответ площадь круга :")
	S = int(P * 2 * R)
	fmt.Println(S)
}
