package main

import "fmt"

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

	answer, err := div(10, 0, "asdasds")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer)
	}
	foo()
	fmt.Println("Hello world!")
}
