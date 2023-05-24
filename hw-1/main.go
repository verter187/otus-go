package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode/utf8"
)

func Dedupe[T comparable](arr []T) map[T]int {
	m := make(map[T]int)
	for _, v := range arr {
		m[v]++
	}
	return m
}

func main() {
	// для более мелких файлов
	filename := "./anna1.txt" //os.Args[1] //"./temp/anna1.txt"
	fContent, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	replace_marks_str := "!,.?:;\"-[]()123456789"
	replace_marks := strings.SplitAfter(replace_marks_str, "")
	content := string(fContent)

	for _, mark := range replace_marks {
		content = strings.ReplaceAll(content, mark, "")
	}

	words := strings.Fields(content)

	fmt.Println(len(words))
	mWords := Dedupe(words)

	names := make([]string, 0, len(mWords))
	for name := range mWords {

		if utf8.RuneCountInString(strings.Trim(name, " ")) < 4 {
			continue
		}

		names = append(names, name)
	}

	sort.Slice(names, func(i, j int) bool {
		return mWords[names[i]] > mWords[names[j]]
	})

	for i, name := range names {
		if i > 9 {
			break
		}
		fmt.Printf("%-7v %v\n", name, mWords[name])
	}
}
