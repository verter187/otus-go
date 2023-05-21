package main

import (
	"fmt"
)

func sum(a, b int) int {
	return a + b
}
func main() {
	fmt.Println(sum(5, 7))

	//Не инициализированный map равен nil
	var c map[int]string
	fmt.Println(c == nil)

	//Инициализированный map не равен nil
	var c1 map[int]string = make(map[int]string)
	fmt.Println(c1 == nil)

	//Не инициализированный slice равен nil
	var s []string
	fmt.Println(s == nil)

	//Инициализированный slice не равен nil
	var s1 []string = []string{}
	fmt.Println(s1 == nil)
}
