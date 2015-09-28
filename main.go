package main

import (
	"encoding/json"
	"fmt"
	"github.com/kellydunn/golang-geo"
	"github.com/simultechnology/myGoSamples/gosample"
	"github.com/simultechnology/myGoSamples/models"
	"github.com/simultechnology/myGoSamples/utils"
	"log"
)

func main() {
	fmt.Println(gosample.StartMessage)

	// Make a few points
	p := geo.NewPoint(42.25, 120.2)
	p2 := geo.NewPoint(30.25, 112.2)

	// find the great circle distance between them
	dist := p.GreatCircleDistance(p2)
	fmt.Printf("great circle distance: %d\n", dist)

	sum := utils.Sum(12, 67, 6)
	fmt.Printf("sum : %d\n", sum)

	person := &models.Person{
		ID:      1,
		Name:    "John",
		Email:   "test.com",
		Age:     34,
		Address: "hoge hoge",
		Memo:    "aaaa",
	}
	b, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
