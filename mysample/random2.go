package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main()  {

	start := time.Now() / (1000 * 1000)
	fmt.Println(":Time:", start)

	var slice = make(map[int]bool)
	count := 17000000

	for i := 0; len(slice) < count; i++ {
		n := rand.Intn(count)
			slice[n] = true
	}
	//fmt.Print(slice)

	i := 0
	for _, _ = range slice {
		i++
//		if i % 17 == 0 {
//			//fmt.Println(k)
//		} else {
		//fmt.Print(k, " ")
//		}
	}
	fmt.Println("")

	end := time.Now() / (1000 * 1000)
	fmt.Println(":Time:", end)

}
