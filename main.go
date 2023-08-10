package main

import (
	"fmt"
	"math/rand"
)

type MyType string

func (mt MyType) MethodMyType() {
	fmt.Println("Hello from method")
}

func foo() {
	//panic("PANIC")
	fmt.Println("foo")
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

	const (
		Grey = iota
		Black
		White
		Red
	)

	println(Email)
	var mtT MyType = "1233"
	rnd, err := packageFunc(mtT)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(rnd)
	}

	fmt.Println(Grey, Black, White, Red)

	answer, err := div(10, 2, "asdasds")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer)
	}
	foo()
	fmt.Println("Hello world!")

	var mt MyType
	mt.MethodMyType()

	var fruit Fruit

	fruit = "Apple"
	name := Name(fruit)
	fmt.Println(name)

	var str string
	str = "Hello, world!"
	println(string(str[0]))

	bornYear := getRandIntInRange()
	helloDudes(bornYear)
}

func getRandIntInRange() int {
	minV, maxV := 1960, 2020

	return rand.Intn(maxV-minV) + minV
}

type Name string
type Fruit string
