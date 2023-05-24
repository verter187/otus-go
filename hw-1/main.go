package main

import "fmt"

func main() {

	// ИСПОЛЬЗОВАНИЕ ZIRO VALUES

	//  Для слайсов и словарей, zero value - это nil.
	// С таким значением будут работать функции и операции, читающие данные, например:
	key := "test"
	var seq []string            // nil
	var cache map[string]string // nil
	l := len(seq)               // 0
	c := cap(seq)               // 0
	l1 := len(cache)            // 0
	v, ok := cache[key]         // "", false

	fmt.Println(l, l1, c, v, ok)
	// Для слайсов будут так же работать append:
	var seq1 []string            // nil
	seq1 = append(seq1, "hello") // []string{"hello"}

	// УДОБНОЕ ИСПОЛЬЗОВАНИЕ, НАПРИМЕРМ ДЛЯ СЛОВАРЯ СЛАЙСОВ:

	type User struct {
		Name string
		Host string
	}

	user1 := User{
		Name: "Jhon",
		Host: "com",
	}

	user2 := User{
		Name: "Bob",
		Host: "ru",
	}

	users := []User{user1, user2}

	// ВМЕСТО:
	// hostUsers := map[string][]string{}
	// for _, user := range users {
	// 	if _, ok := hostUsers[user.Host]; !ok {
	// 		hostUsers[user.Host] = make([]string)
	// 	}
	// 	hostUsers[user.Host] = append(hostUsers[user.Host], user.Name)
	// }

	// МОЖНО:
	hostUsers1 := map[string][]string{}
	for _, user := range users {
		hostUsers[user.Host] = append(hostUsers[user.Host], user.Name)
	}

	fmt.Println(hostUsers1)
}
