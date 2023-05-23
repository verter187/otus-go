package main

import (
	"fmt"
	"reflect"
)

func Concat(slices [][]int) []int {
	var res []int
	for _, s := range slices {
		res = append(res, s...)
	}
	return res
}

func main() {
	type pair struct {
		s [][]int
		r []int
	}
	test := []pair{
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{[][]int{{1, 2}, {3, 4}, {6, 5}}, []int{1, 2, 3, 4, 6, 5}},
		{[][]int{{1, 2}, {}, {6, 5}}, []int{1, 2, 6, 5}},
	}
	for _, t := range test {
		s := t.s
		r := t.r
		r2 := Concat(s)
		fmt.Printf("Test: %v\n", s)
		fmt.Printf("Expected: %v\n", r)
		fmt.Printf("Result: %v\n", r2)
		if reflect.DeepEqual(r, r2) {
			fmt.Printf("OK\n\n")
		} else {
			fmt.Printf("FAIL\n\n")
		}
	}
}
