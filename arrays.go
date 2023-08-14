package main

import "fmt"

//var lastWeekTemp [7]int

/*
Когда вместо range лучше использовать цикл for?

	Когда требуется изменить порядок обхода
	Когда не требуется обход всех элементов
*/
func arraysSimple() {
	rgbColor := [...]uint8{255, 255, 128, 1}
	fmt.Println(rgbColor)

	weekTemp := [7]int{5, 4, 6, 8, 11, 9, 5}
	sumTemp := 0

	for _, temp := range weekTemp {
		sumTemp += temp
	}

	average := sumTemp / len(weekTemp)
	fmt.Println(average)
}
