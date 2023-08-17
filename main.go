package main

import (
	"bufio"
	"fmt"
	"goFirstTry/toppackage/middlepackage/bottompackage/mathxxx"
	"math/rand"
	"os"
)

func foo() {
	panic("PANIC")
}

func div(a, b int, x string) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("делитель равен 0")
	}

	fmt.Println(x)
	return a / b, nil
}

func main() {

	// Срабатывает при завершении main
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	helloDudes(getRandIntInRange())
	// операторы ветвления
	simpleIf("апыварывпо")
	switchCaseSimple(12)
	switchCaseFallThrough("hello")
	switchCaseFallThrough("world")
	// циклы
	simpleFor()
	infiniteLoop()
	inRangeLoop()
	fizzBuzz()

	printPointer()
	exampleUsePointers()
	//rowsCounter()
	arraysSimple()
	simpleSlices()

	//fillSliceExample()
	//maps()
	//order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	//mapsTest(order)

	//indexes := find([]int{3, 1, 4, 10, 2, 14, 10, 12, 13, 1}, 4)
	//fmt.Println(indexes)

	//input := []string{
	//	"cat",
	//	"dog",
	//	"bird",
	//	"dog",
	//	"parrot",
	//	"cat",
	//}
	//
	//fmt.Println(RemoveDuplicates(input))

	//var user, err = NewUser("alex", "email@email.com", 24)
	//
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(*user)
	//}
	//
	//jsonUser, err := json.Marshal(*user)
	//if err != nil {
	//	log.Fatalln("error")
	//}
	//
	//fmt.Println(string(jsonUser))
	funcArea, ok := area(square)

	if !ok {
		fmt.Println("Неизвестная фигура")
	} else {
		s := funcArea(5)
		fmt.Println(s)
	}

	//VeryLongTimeFunction()
	//testDefer()

	//var s = []int{1, 2, 4, 31, 13, 12}
	//
	//fmt.Println(s)
	//
	//sum := mathSlice.SumSlice(s)
	//fmt.Println(sum)
	//
	//mathSlice.MapSlice(s, func(el int) int {
	//	return el * el
	//})
	//
	//fmt.Println(s)
	//
	//result := mathSlice.FoldSlice(s, func(x, y int) int {
	//
	//	if y%2 == 0 {
	//		return x + y
	//	}
	//	return x
	//}, 0)
	//
	//fmt.Println(result)

	fmt.Println(mathxxx.AddInts(4, 7))

}

func getRandIntInRange() int {
	minV, maxV := 1960, 2020
	return rand.Intn(maxV-minV) + minV
}

func rowsCounter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	cnt := 0
	for {
		fmt.Print("-> ")
		// Считываем введённую пользователем строку. Программа ждёт, пока пользователь введёт строку
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		f(&cnt)

		fmt.Printf("User input %d lines\n", cnt)
	}
}

func f(pCnt *int) {
	*pCnt++
}
