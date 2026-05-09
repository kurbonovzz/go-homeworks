package main

import "fmt"

var products []string

func addProduct(name string) {
	products = append(products, name)
}

func showProducts() {
	fmt.Println("Список покупок")
	fmt.Println(products)
}

func main() {

	var choice int
	var name string

	for {
		fmt.Println("\n1 - Добавить продукт")
		fmt.Println("2 - Показать список")
		fmt.Println("0 - Выход")

		fmt.Print("Выбор: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			fmt.Print("Введите продукт: ")
			fmt.Scan(&name)

			addProduct(name)

		case 2:
			showProducts()

		case 0:
			return

		default:
			fmt.Println("Ошибка")
		}
	}
}
