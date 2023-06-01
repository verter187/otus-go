package main

import (
	"fmt"
)

// области видимости, блоки
// функции: вариадические, анонимные, замыкания, методы
// ошибки: принципы обработки, best practices
// panic, recover, defer

// Области видимости и блоки
// Блок (block) - это, возможно пустая, последовательность объявлений (declarations) и операторов (statements) в соответствующих фигурных скобках

// - universe block - весь код проекта
// - package block - весь код пакета
// - file block - исходный код в файле
// - local block - просто {}
func main() {
	{ // yначало блока 1
		a := 1
		fmt.Println(a)
		{ // начало блока 2
			b := 2
			fmt.Println(b)
		} // конец блока 2
	} // конец блока 1

	// Неявные блоки: for, if, switch, case, select

	// {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//}

	if i := 0; i >= 0 {
		fmt.Println(i)
	}

	switch i := 2; i * 4 {
	case 8:
		j := 0
		fmt.Println(i, j)
	default:
		// "j" is undefined here
		fmt.Println("default")
	} // "j" is undefined here

	// Интересное:
	// * Нет дефолтных значений для параметров
	// * функция может возвращать несколько значений
	// * функция - first class value, можем работать как с обычным значением
	// * параметры в функцию передаются по значению

	func(text string) {
		fmt.Println(text)
	}("test777")

	sum := func(a ...int) int {

		s := 0
		for _, i := range a {
			s += i
		}
		return s
	}
	s := sum(1, 2, 3, 4, 5)
	fmt.Println(s)

	p := []int{1, 2, 3, 4, 5, 6}
	s1 := sum(p...)
	fmt.Println(s1)

	// Замыкания

	intSeq := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
