package main

import "fmt"

func main()  {

	slice := make([]int, 0 ,10)

	fmt.Println(slice)
	slice = append(slice, 10)
	fmt.Println(slice)


	slice2 := make([]int, 10)

	fmt.Println(slice2)
	fmt.Println(slice2[7])

}
