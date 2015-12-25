package main

import (
	"fmt"
)

type Saiyan struct {
	Name string
	Friends map[string]*Saiyan
}

func main()  {
	fmt.Println("start!")

	lookup := make(map[string]int)
	lookup["goku"] = 9001
	//lookup["goku2"] = 9002

	power, exist := lookup["vegeta"]

	fmt.Println(power, exist)

	length := len(lookup)
	fmt.Printf("len : %d\n", length)

	delete(lookup, "goku")

	length = len(lookup)
	fmt.Printf("len : %d\n", length)

}
