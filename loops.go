package main

import (
	"fmt"
	"math/rand"
)

func simpleFor() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func infiniteLoop() {
	fmt.Println("Start infinite loop")
	for {
		x := rand.Intn(10)
		fmt.Println(x)
		if x > 8 {
			break
		}
	}
	fmt.Println("End infinite loop")
}

func inRangeLoop() {
	array := [3]int{1, 2, 3}

	for index, value := range array {
		fmt.Printf("array[%d]: %d\n", index, value)
	}
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {

		if i%15 == 0 {
			fmt.Println("FizzBazz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("Bazz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}

		fmt.Println(i)
	}
}
