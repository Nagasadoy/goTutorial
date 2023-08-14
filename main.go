package main

import (
	"bufio"
	"fmt"
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
