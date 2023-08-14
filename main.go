package main

import (
	"fmt"
	"math/rand"
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
}

func getRandIntInRange() int {
	minV, maxV := 1960, 2020
	return rand.Intn(maxV-minV) + minV
}
