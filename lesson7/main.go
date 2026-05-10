package main

import "fmt"

type Tasks interface {
	printAllTasks()
	addWorkTaskDay()
	addWeekendTask()
	printWorkTaksDay()
	printWeekendTask()
}

func printAllTasks(tasks []string) {
	fmt.Println("\n------------Все задачи-------------- ")
	for i := 0; i < len(tasks); i++ {
		fmt.Println(tasks[i])
	}
}

func addWorkTaskDay(workTaskDay []string) {
	var tasks string
	fmt.Println("\n------Новая задача для Понедельник------------- ")
	fmt.Scan(&tasks)

	workTaskDay[0] = "Понедельник " + tasks
}

func addWeekendTask(weekend []string) {
	var tasks string
	fmt.Println("\n-------Новая задача для суббота-----------")
	fmt.Scan(&tasks)

	weekend[0] = "Суббота " + tasks

}

func printWorkTakesDay(workDays []string) {
	fmt.Println("\n--------------Будниие дни----------")
	for _, tasks := range workDays {
		fmt.Println(tasks)
	}
}

func printWeekendTask(weekend []string) {
	fmt.Println("\n---------Выходные---------")
	for _, tasks := range weekend {
		fmt.Println(tasks)
	}
}

func main() {

	tasks := []string{
		"понедельник: учеба ",
		"Вторник: спорт",
		"Среда: фылм",
		"Четверг: футбол",
		"Пятница: репетиция",
		"Суббота: прогулка",
		"Воскресные: сон",
	}
	workDays := tasks[0:5]
	weekend := tasks[5:7]

	printAllTasks(tasks)
	printWorkTakesDay(workDays)
	printWeekendTask(weekend)
	addWorkTaskDay(workDays)
	addWeekendTask(weekend)
	printAllTasks(tasks)

}
