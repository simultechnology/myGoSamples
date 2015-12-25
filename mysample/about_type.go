package main

import (
	"fmt"
	"reflect"
	"math"
)

func main()  {

	fmt.Println("start")
	fmt.Println(reflect.TypeOf(16))
	n := math.Pow(2, 449)
	fmt.Println(n)
	fmt.Println(reflect.TypeOf(n))
}
