package main

import (
	"fmt"
)

func main()  {
	fmt.Println("start");

	name := "ishi"
	greeting := fmt.Sprintf("my name is %s", name)

	greeting += " ! hoge"
	fmt.Println(greeting)

	var firstName string

	fmt.Println(len(firstName))
}
