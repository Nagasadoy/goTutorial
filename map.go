package main

import "fmt"

func maps() {
	m := make(map[string]string)
	m["Foo"] = "foo"
	m["Bar"] = "bar"

	v, ok := m["S"] // в s запишется (true - если значение по ключу есть, false - если значения по ключу нет,
	// при этом в v запишется дефолтное значение для этого типа)

	fmt.Println(v)
	fmt.Println(ok)

	fmt.Println(m)

	// Функции для работы с map

	mapLength := len(m)

	fmt.Println(mapLength)
	// сравнивать map можно только с nil а между собой нельзя
	//var m map[string]string
	//if m != nil {            // если не проверить это условие,
	//	m["foo"] = "bar"    // то здесь можно получить panic
	//}

	delete(m, "Bar")
	fmt.Println(m)

	m["bar"] = "b ar"
	m["cool"] = "fd"
	m["cat"] = "meow"

	for k, v := range m {
		fmt.Printf("Ключ %v имеет значение %v\n", k, v)
	}

	/*
		Если сделать m := make(map[int]int), не заполнить данными и всё же запросить значение ключа 100 v :=  m[100],
		запрос будет отработан и вернёт значение 0 (нулевое значение для типа int).
		Если присвоить значение ключу 50 m[50] = 0 и запросить его v := m[50], ответ будет таким же — 0.
		Это два разных случая:
		ключу не назначено значение;
		ключу назначено нулевое значение.
		Чтобы различать их, лучше пользоваться полной формой индексного выражения: v, ok = m[k].
		Тогда переменная ok примет значение true, если ключ найден, и false в противном случае.
	*/
}

func mapsTest(order []string) {

	products := map[string]int{
		"хлеб": 50, "молоко": 100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}

	for k, v := range products {
		if v > 500 {
			fmt.Println(k)
		}
	}

	sumTotal := 0
	for _, item := range order {
		sumTotal += products[item]
	}

	fmt.Println(sumTotal)
}

func mapTest2(a []int, k int) []int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == k {
				return []int{i, j}
			}
		}
	}

	return make([]int, 0)
}

func find(arr []int, k int) []int {
	// Создадим пустую map
	m := make(map[int]int)
	// будем складывать в неё индексы массива, а в качестве ключей использовать само значение
	for i, a := range arr {
		if j, ok := m[k-a]; ok { // если значение k-a уже есть в массиве, значит, arr[j] + arr[i] = k и мы нашли, то что нужно
			return []int{i, j}
		}
		// если искомого значения нет, то добавляем текущий индекс и значение в map
		m[a] = i
	}
	// не нашли пары подходящих чисел
	return nil
}

func RemoveDuplicates(input []string) []string {
	duplicates := make(map[string]string)
	noDuplicates := make([]string, 0)

	for _, v := range input {
		if _, ok := duplicates[v]; !ok {
			noDuplicates = append(noDuplicates, v)
		}
		duplicates[v] = v
	}

	return noDuplicates
}
