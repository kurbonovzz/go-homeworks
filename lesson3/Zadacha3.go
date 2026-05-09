package main

import "fmt"

func zadacha8() {

	fmt.Println("Введите два  числа :")

	var a int
	fmt.Scan(&a)

	var b int
	fmt.Scan(&b)

	fmt.Println("Найдите их среднее арифметическое :")

	var C int

	C = (a + b) / 2
	fmt.Println(C)

}
