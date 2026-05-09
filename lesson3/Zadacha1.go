package main

import "fmt"

func zadacha1() {

	fmt.Println("Введите прямоугольного параллелепипеда:")

	var a, b, c int
	fmt.Scan(&a, &b, &c)

	fmt.Println("Найти его и площадь поверхности: ")

	V := (a * b * c)
	S := 2 * (a*b + b*c + a*c)
	fmt.Println(V, S)
}
