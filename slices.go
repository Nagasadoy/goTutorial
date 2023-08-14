package main

import (
	"fmt"
	"reflect"
)

func simpleSlices() {
	//var mySlice []int // по умолчанию nil
	//fmt.Println(mySlice)

	//mySlice := make([]TypeOfelement, LenOfslice, CapOfSlice)
	mySlice := make([]int, 0)    // слайс [], базовый массив []
	mySlice = make([]int, 5)     // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0]
	mySlice = make([]int, 5, 10) // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0 0 0 0 0 0]

	fmt.Println(mySlice)

	weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
	workDaysSlice := weekTempArr[:5]
	weekendSlice := weekTempArr[5:]
	fromTuesdayToThursDaySlice := weekTempArr[1:4]
	weekTempSlice := weekTempArr[:]

	fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice))                                        // [1 2 3 4 5] 5 7
	fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice))                                           // [6 7] 2 2
	fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6
	fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))                                        // [1 2 3 4 5 6 7] 7 7

	// Для добавления используется функция append

	a := make([]int, 1000)
	b := append(a, 7)
	fmt.Println(len(a), cap(a)) // 1000 1000
	fmt.Println(len(b), cap(b)) // 1001 1536

	s1 := []int{1, 2, 3, 4, 5, 10, 112}
	s2 := []int{12, 32, 44, 1123}
	s1 = append(s1[:], s2...) // объединение 2 слайсов
	fmt.Println(s1)

	// для копирования dest

	var dest []int
	dest2, dest3 := make([]int, 3), make([]int, 5)
	src := []int{1, 2, 3, 4}
	copy(dest, src)
	copy(dest2, src)
	copy(dest3, src)
	fmt.Println(dest, dest2, dest3, src)

	exampleSlice := []int{1, 2, 4, 2, 2312, 123, 231, 43}
	exampleSlice2 := []int{1, 2, 4, 2, 2312, 123, 231, 1}

	exampleSlice = deleteLastElement(exampleSlice)
	fmt.Println(exampleSlice)
	exampleSlice = deleteFirstElement(exampleSlice)
	fmt.Println(exampleSlice)
	exampleSlice = deleteElementI(exampleSlice, 3)
	fmt.Println(exampleSlice)

	sliceIsEquals(exampleSlice, exampleSlice2)
}

func deleteLastElement(slice []int) []int {
	if len(slice) != 0 {
		s := slice[:len(slice)-1]
		return s
	}
	return slice
}

func deleteFirstElement(slice []int) []int {
	if len(slice) != 0 {
		s := slice[1:]
		return s
	}
	return slice
}

func deleteElementI(slice []int, i int) []int {
	if len(slice) != 0 && i < len(slice) { // защищаемся от паники
		return append(slice[:i], slice[i+1:]...)
	}

	return slice
}

func sliceIsEquals(slice1 []int, slice2 []int) bool {
	return reflect.DeepEqual(slice1, slice2)
}

func fillSliceExample() {

	dim := 100
	slice := make([]int, 0, dim)

	for i := 0; i < dim; i++ {
		slice = append(slice, i+1)
	}

	slice = append(slice[:10], slice[dim-10:]...)

	for i := range slice[:dim/2] {
		slice[i], slice[dim-i-1] = slice[dim-i-1], slice[i]
	}

	fmt.Println(slice)
}
