package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	fmt.Println("start!")

	file, err := os.Create("./tmp/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	for n := 0; n < 1000; n++ {
		file.WriteString(fmt.Sprintf("num = %d\n", n))
	}
	file.WriteString("end!")

	fmt.Println("end!")
}
