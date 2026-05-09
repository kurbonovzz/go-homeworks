package main

import "fmt"

func PutGrade(grades *[5]int, grade int, index int) {
	if grade > 10 {
		return
	}

	grades[index] = grade
}

func showGrades(grades [5]int) {
	fmt.Println("Оценки студента ", grades)

	sum := 0

	for i := 0; i < 5; i++ {
		sum += grades[i]
	}

	fmt.Println("Сумма оценок", sum)
	fmt.Println("Средня оценка", float64(sum)/5)
}

func main() {
	var grades [5]int

	PutGrade(&grades, 1, 0)
	PutGrade(&grades, 4, 1)
	PutGrade(&grades, 3, 2)
	PutGrade(&grades, 5, 3)
	PutGrade(&grades, 4, 4)

	showGrades(grades)
}
