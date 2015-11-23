package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
	*Person
	Power int
}

func (s *Saiyan) Introduce() {
	fmt.Printf("Hi, I'm %s ya! \n", s.Name)
}

func main()  {
	fmt.Println("start!")

	taka := &Person{"taka"}
	taka.Introduce()

	goku := &Saiyan{
		&Person{"goku"},
		9001,
	}

	goku.Introduce()
	goku.Person.Introduce()
}

