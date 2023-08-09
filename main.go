package main

import "fmt"

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
}

type Name string
type Fruit string
