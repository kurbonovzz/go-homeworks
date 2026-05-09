package main

import "fmt"

func zadacha3() {

	fmt.Println("Даны стороны прямоугольника :")
	var a int
	fmt.Scan(&a)

	var b int
	fmt.Scan(&b)

	fmt.Println("Найти его площадь и периметр:")
	P := (4 * a)
	S := (a * a)

	fmt.Println(P, S)

}
