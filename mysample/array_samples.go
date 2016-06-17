package main

import (
	"fmt"
)

func main() {
	fmt.Println("start!")

	name := []string{"ishi"}

	printName(name)

}

func printName(names []string)  {

	fmt.Printf("%v\n", names)

}
