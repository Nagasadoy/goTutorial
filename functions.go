package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"
)

type figures int

const (
	square figures = iota
	circle
	triangle
)

func area(figures figures) (func(float64) float64, bool) {

	pi := 3.14

	switch figures {
	case square:
		return func(x float64) float64 {
			return x * x
		}, true
	case circle:
		return func(r float64) float64 {
			return pi * r * r
		}, true
	case triangle:
		return func(a float64) float64 {
			return a * a * math.Sqrt(3) / 4
		}, true
	default:
		return nil, false
	}
}

func PrintAllFilesWithFuncFilter(path string, predicate func(string) bool) {
	// создаём переменную, содержащую функцию обхода
	// мы создаём её заранее, а не через оператор :=, чтобы замыкание могло сослаться на него
	var walk func(string)
	walk = func(path string) {
		// получаем список всех элементов в папке (и файлов, и директорий)
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}
		//  проходим по списку
		for _, f := range files {
			// получаем имя элемента
			// filepath.Join — функция, которая собирает путь к элементу с разделителями
			filename := filepath.Join(path, f.Name())
			// печатаем имя элемента, если predicate вернёт true
			if predicate(filename) {
				fmt.Println(filename)
			}
			// если элемент — директория, то вызываем для него рекурсивно ту же функцию
			if f.IsDir() {
				walk(filename)
			}
		}
	}
	// теперь вызовем функцию walk
	walk(path)
}

func metricTime(start time.Time) {
	// функция Now() возвращает текущее время, а функция Sub возвращает разницу между двумя временными метками
	fmt.Println(time.Now().Sub(start))
}

func VeryLongTimeFunction() {
	defer metricTime(time.Now()) // передаём в функцию metricTime значение текущего времени и откладываем её вызов до возврата
	for i := 0; i < 4; i++ {
		time.Sleep(time.Second)
	}
}

func testDefer() {
	a := "goodbye"
	defer func() {
		fmt.Println(a)
	}()

	fmt.Println("hello")
}
