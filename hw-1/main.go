package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func strUnpack(str string) (string, error) {
	var ns string
	prev := ""
	for _, s := range str {

		cs := string(s)
		if num, err := strconv.Atoi(string(s)); err == nil {
			if prev == "" || isNumeric(prev) {
				return "", errors.New("invalid string")
			} else {
				cs = strings.Repeat(prev, num-1) //already printed once
			}
		}

		prev = string(s)
		ns += cs
	}

	return ns, nil
}

func main() {

	// * "a4bc2d5e" => "aaaabccddddde"
	// * "abcd" => "abcd"
	// * "45" => "" (некорректная строка)
	if test, err := strUnpack("45"); err == nil {
		fmt.Println(test)
	} else {
		fmt.Println(err)
	}
}
