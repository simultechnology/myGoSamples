package main

import (
	"fmt"
)

type User struct  {
	name string
	age int
}
func (u *User) addAge(age int) {
	u.age += age
	fmt.Printf("age : %d\n", u.age)
}

func main()  {
	fmt.Println("start!")

	var u User
	u.name = "taka"
	u.age = 38
	fmt.Printf("name: %s, age: %d\n", u.name, u.age)

	u.addAge(10)

	fmt.Printf("name: %s, age: %d\n", u.name, u.age)

	var ken *User = &User{name: "ken", age: 30}
	fmt.Printf("name: %s, age: %d\n", ken.name, ken.age)
	ken.addAge(5)
	fmt.Printf("name: %s, age: %d\n", ken.name, ken.age)

}
