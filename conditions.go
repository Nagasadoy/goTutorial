package main

import "fmt"

func simpleIf(str string) {
	if len(str) > MaxLength {
		fmt.Println(MaxLengthMessage)
	} else {
		fmt.Println("Фигня")
	}
}
