package main

import "fmt"

type Employee struct {
	name, surname string
}

func FullName(e Employee) string {
	return e.name + " " + e.surname
}

func (e Employee) FullName() string {
	return e.name + " " + e.surname
}

func main() {
	fmt.Println(Employee{"alexander", "davydov"}.FullName())
	fmt.Println(FullName(Employee{"alexander", "davydov"}))
}
