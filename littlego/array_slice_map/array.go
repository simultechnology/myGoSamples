package main

import (
	"fmt"
)

func main()  {

	arr := [3]int{4, 5, 9}
	fmt.Println(arr)

	var arr2 [10]int

	fmt.Println(arr2)
	arr2[7] = 90
	fmt.Println(arr2)

	// error
	//append := append(arr2, 765)

	slice := []int{56, 67, 66}
	fmt.Println(slice)
	slice = append(slice, 8888)
	fmt.Println(slice)

	i := make([]int, 10)
	fmt.Println(i)

}


