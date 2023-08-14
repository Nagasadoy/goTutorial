package main

import (
	"fmt"
	"time"
)

func printPointer() {
	var a int = 5
	p := &a
	fmt.Println(a, p) //a=5 p=0xc0000b2008
}

type A struct {
	IntField int
}

func pointerCommons() {

	p := &A{
		IntField: 10,
	}

	fmt.Println(p)

	//p = new(A) // другой способ записи &A{}

	// неявное разыменование при доступе к полям структуры
	p.IntField = 42

	// как в C можно делать указатели на указатели **int

	var pp *int
	fmt.Println(*pp) // panic: runtime error: invalid memory address or nil pointer dereference
}

type Person struct {
	Name        string
	Age         int
	lastVisited time.Time
}

func UpdatePersonWithLastVisited(p *Person) {
	p.lastVisited = time.Now()
}

func exampleUsePointers() {
	p := Person{
		Name: "ALEX",
		Age:  25,
	}

	UpdatePersonWithLastVisited(&p)
}
