package main

import (
	"fmt"
	"math/rand"
)
func main()  {

	var slice = make(map[int]bool)

	for i := 0; len(slice) < 1700000; i++ {
		n := rand.Intn(1700000)
		if !slice[n] {
			slice[n] = true
		}
	}
	//fmt.Print(slice)

	i := 0
	for k, _ := range slice {
		i++
		if i % 17 == 0 {
			fmt.Println(k)
		} else {
			fmt.Print(k, " ")
		}
	}

}
