package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
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

	var user, err = NewUser("alex", "email@email.com", 24)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*user)
	}

	jsonUser, err := json.Marshal(*user)
	if err != nil {
		log.Fatalln("error")
	}

	fmt.Println(string(jsonUser))
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
