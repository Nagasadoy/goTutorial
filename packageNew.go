package main

import (
	"fmt"
	"math/rand"
)

const Email = "nagasadoy@gmail.com"

func packageFunc(myType MyType) (int, error) {
	if len(myType) < 2 {
		return 0, fmt.Errorf("длина строки меньше 2")
	}

	return rand.Int(), nil

}

func helloDudes(year int) {
	fmt.Println(year)
	switch {
	case year > 1946 && year <= 1964:
		fmt.Println("Hello boomer")
	case year > 1965 && year <= 1980:
		fmt.Println("Hello X")
	case year > 1981 && year <= 1996:
		fmt.Println("Hello Millenial")
	case year > 1997 && year <= 2012:
		fmt.Println("Hello zoomer")
	case year > 2013:
		fmt.Println("Hello alfa")
	default:
		fmt.Println("WTF")
	}
}
