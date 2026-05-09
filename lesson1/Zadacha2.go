package main

import "fmt"

func zadacha2() {

	fmt.Println("Найдите периметр прямоугольника ")

	fmt.Println("Введите а")
	var a int
	fmt.Scan(&a)

	fmt.Println("Введите b")
	var b int
	fmt.Scan(&b)

	fmt.Println("Ответ:")
	fmt.Println(2 * (a + b))

}
