package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
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

type Stopwatch struct {
	startValue time.Time
	stopValues []time.Duration
}

func (sw *Stopwatch) Start() {
	sw.startValue = time.Now()
	sw.stopValues = nil
}

func (sw *Stopwatch) SaveSplit() {
	now := time.Now()
	delta := now.Sub(sw.startValue)
	sw.stopValues = append(sw.stopValues, delta)
}

func (sw *Stopwatch) GetResults() []time.Duration {
	return sw.stopValues
}

////////////////////////////////////////////////////////////////////

type Person struct {
	Name string
	Year int
}

func NewPerson(name string, year int) Person {
	return Person{
		Name: name,
		Year: year,
	}
}

func (p Person) String() string {
	return fmt.Sprintf("Имя: %s, Год рождения: %d", p.Name, p.Year)
}

func (p Person) Print() {
	fmt.Println(p) // вызовется метод String
}

type Student struct {
	Person // вложенный объект Person
	Group  string
}

func NewStudent(name string, year int, group string) Student {
	return Student{
		Person: NewPerson(name, year), // Явно создаём структуру Person
		Group:  group,
	}
}

// String возвращает информацию о студенте.
func (s Student) String() string {
	return fmt.Sprintf("%s, Группа: %s", s.Person, s.Group)
}

func (s Student) Print() {
	fmt.Println(s) // вызовется метод String
}

func (s *Student) Debug() {
	// доступ к методам объекта Person
	s.Print()
	// или
	s.Person.Print()

	// доступ к полю 'Name' объекта Person
	s.Name = "Mark Smith"
	// или
	s.Person.Name = "Mark Smith"

	// вызовется метод String объекта Student
	fmt.Println(s)
	// вызовется метод String объекта Person
	fmt.Println(s.Person)
}

func main() {

	// Срабатывает при завершении main
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	//helloDudes(getRandIntInRange())
	//// операторы ветвления
	//simpleIf("апыварывпо")
	//switchCaseSimple(12)
	//switchCaseFallThrough("hello")
	//switchCaseFallThrough("world")
	//// циклы
	//simpleFor()
	//infiniteLoop()
	//inRangeLoop()
	//fizzBuzz()
	//
	//printPointer()
	//exampleUsePointers()
	////rowsCounter()
	//arraysSimple()
	//simpleSlices()
	//
	////fillSliceExample()
	////maps()
	////order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	////mapsTest(order)
	//
	////indexes := find([]int{3, 1, 4, 10, 2, 14, 10, 12, 13, 1}, 4)
	////fmt.Println(indexes)
	//
	////input := []string{
	////	"cat",
	////	"dog",
	////	"bird",
	////	"dog",
	////	"parrot",
	////	"cat",
	////}
	////
	////fmt.Println(RemoveDuplicates(input))
	//
	////var user, err = NewUser("alex", "email@email.com", 24)
	////
	////if err != nil {
	////	fmt.Println(err)
	////} else {
	////	fmt.Println(*user)
	////}
	////
	////jsonUser, err := json.Marshal(*user)
	////if err != nil {
	////	log.Fatalln("error")
	////}
	////
	////fmt.Println(string(jsonUser))
	//funcArea, ok := area(square)
	//
	//if !ok {
	//	fmt.Println("Неизвестная фигура")
	//} else {
	//	s := funcArea(5)
	//	fmt.Println(s)
	//}
	//
	////VeryLongTimeFunction()
	////testDefer()
	//
	////var s = []int{1, 2, 4, 31, 13, 12}
	////
	////fmt.Println(s)
	////
	////sum := mathSlice.SumSlice(s)
	////fmt.Println(sum)
	////
	////mathSlice.MapSlice(s, func(el int) int {
	////	return el * el
	////})
	////
	////fmt.Println(s)
	////
	////result := mathSlice.FoldSlice(s, func(x, y int) int {
	////
	////	if y%2 == 0 {
	////		return x + y
	////	}
	////	return x
	////}, 0)
	////
	////fmt.Println(result)
	//
	//fmt.Println(mathxxx.AddInts(4, 7))

	//sw := Stopwatch{}
	//sw.Start()
	//time.Sleep(1 * time.Second)
	//sw.SaveSplit()
	//
	//time.Sleep(500 * time.Millisecond)
	//sw.SaveSplit()
	//
	//time.Sleep(300 * time.Millisecond)
	//sw.SaveSplit()
	//
	//fmt.Println(sw.GetResults())

	student := NewStudent("Igor", 23, "math")
	student.Print()
	fmt.Println(student)

}

func getRandIntInRange() int {
	minV, maxV := 1960, 2020
	return rand.Intn(maxV-minV) + minV
}

func rowsCounter() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	cnt := 0
	for {
		fmt.Print("-> ")
		// Считываем введённую пользователем строку. Программа ждёт, пока пользователь введёт строку
		_, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		f(&cnt)

		fmt.Printf("User input %d lines\n", cnt)
	}
}

func f(pCnt *int) {
	*pCnt++
}

type ExtendedLogger struct {
	log.Logger
	LogLevel
}

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)
