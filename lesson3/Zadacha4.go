package main

import (
	"fmt"
	"math"
)

func zadacha9() {

	fmt.Println("Введите два неотрицательных числа: ")
	fmt.Println("Первое число :")
	var a int
	fmt.Scan(&a)

	fmt.Println("Второе число :")
	var b int
	fmt.Scan(&b)
	fmt.Println("Найти их среднее геометрическое:")

	S := math.Sqrt(float64(a * b))

	fmt.Println(S)

}
