package main

import (
	"fmt"
	"github.com/simultechnology/myGoSamples/utils"
	"os"
	"log"
	"github.com/simultechnology/myGoSamples/gosample"
	"encoding/json"
	"github.com/simultechnology/myGoSamples/models"
)

func main() {
	sum := utils.Sum(6, 9, 6)
	fmt.Printf("sum is %d\n", sum)

	file, err := os.Create("./file.txt")
	printError(err)
	defer file.Close()

	_, err = file.Write([]byte(gosample.Message))
	printError(err)

	file, err = os.Create("./tmp/sample.json")
	printError(err)
	defer file.Close()

	p, err := json.Marshal(&models.Person{ID: 5})
	file.Write(p)

}

func printError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}