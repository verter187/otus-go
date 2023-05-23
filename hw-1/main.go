package main

import "fmt"

func main() {

	var cache map[string]string //не инициализированный словарь, nil

	cache1 := map[string]string{} // созданный с помощью литерала, len(cache) == 0

	cache2 := map[string]string{ // литерал с первоначальным значением
		"one":   "один",
		"two":   "два",
		"three": "три",
	}

	cache3 := make(map[string]string)      // тоже что и map[string]string
	cache4 := make(map[string]string, 100) // заранее выделить память на 100 ключей

	fmt.Println(cache, cache1, cache2, cache3, cache4)

	key := "two" //Получение знаения, если ключ не найден - Zero Value
	value := cache2[key]
	fmt.Println(value)

	empval := cache[key] // если ключ не найден - Zero Value
	fmt.Println(empval == "")

	_, ok := cache2[key] // проверить наличие ключа в словаре
	fmt.Println(ok)

	cache1[key] = value // Записать значение в инициализированный(!) словарь
	fmt.Println(cache1[key])

	delete(cache, key) // удалить ключ из словаря, работает всегда

	//В go нет функций, возвращающих списки ключей и значений словаря. (Почему?)
	// Получить ключи
	var keys []string
	for key, _ := range cache2 {
		keys = append(keys, key)
	}

	fmt.Println(keys)

	// Получить Значения
	values := make([]string, 0, len(cache))
	for _, val := range cache2 {
		values = append(values, val)
	}
	fmt.Println(values)

	// Ключом может быть любой тип данных, для которго определена операция сравнения (==):
	// * строки, числовые типы, булево
	// * каналы (chan)
	// * интерфейсы
	// * указатели
	// * структуры или массивы, содержащие сравнимые типы

	type User struct {
		Name string
		Host string
	}

	var cache5 map[User][]int
	fmt.Println(cache5)

}
