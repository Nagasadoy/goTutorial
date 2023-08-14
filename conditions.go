package main

import "fmt"

func simpleIf(str string) {
	if len(str) > MaxLength {
		fmt.Println(MaxLengthMessage)
	} else {
		fmt.Println("Фигня")
	}
}

func switchCaseSimple(a int) {
	switch b := a % 5; { // можно вот тут объявлять переменные
	case b == 0:
		fmt.Println("Кратно 5")
	default:
		fmt.Printf("Остаток от деления на 5: %d", b)
	}
}

/*
Чтобы досрочно прервать выполнение case, используют ключевое слово break.
Это бывает полезно, когда внутри case есть условные конструкции.
В Go нет необходимости явно указывать break в конце каждого case,
так как следующий блок case автоматически не выполнится при совпадении условия.
Когда нужно всё-таки выполнить следующий блок, используют ключевое слово fallthrough.
Если указать его в конце блока кода, то после него будет выполнен блок в следующем case
или default.

У ключевого слова fallthrough есть особенности:
  - его можно использовать только в последней строке case, иначе будет ошибка компиляции;
  - оно игнорирует условие следующего по порядку case.
*/
func switchCaseFallThrough(value string) {
	switch value {
	case "hello":
		fmt.Println("Hello")
		fallthrough
	case "world":
		fmt.Println("World")
	default:
		fmt.Println("???")
	}
}
