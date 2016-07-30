package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("start!")

	name := []string{"ishi"}

	printName(name)

	defaultKeywords := []string{"aaa","bb","あいうえお"}

	meta := &Meta{
		Keywords: []string{"東京","神田","秋葉原"},
	}

	keywords := append(defaultKeywords, meta.Keywords...)

	fmt.Printf("keywords : %v\n", keywords)

	fmt.Printf("%v\n", []rune("Hello, 世界 XI"))
	newKeywords := make([]string, 10)
	copy(newKeywords, keywords)
	fmt.Printf("%v\n", keywords[:5])
	fmt.Printf("%v\n", newKeywords[:8])
	fmt.Printf("%v\n", len(newKeywords[:8]))

	newKeywords = append(newKeywords, "")
	fmt.Printf("%v\n", newKeywords)

	var realKeywords []string
	for _, word := range newKeywords {
		if len(word) > 0 {
			realKeywords = append(realKeywords, word)
		}
	}

	fmt.Printf("%s\n", strings.Join(newKeywords, ","))
	fmt.Printf("%s\n", strings.Join(realKeywords, ","))

}

type Meta struct {
	Title       string
	Keywords    []string
	Description string
}

func printName(names []string) {

	fmt.Printf("%v\n", names)


}
