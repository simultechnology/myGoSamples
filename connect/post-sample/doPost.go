package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"html/template"
	"io/ioutil"
)

type Person struct {
	ID   int        `json:"id"`
	Name string    `json:"name"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request)  {

	fmt.Printf("Hello world!!!\n")
}

var t = template.Must(template.ParseFiles("index.html"))

func PersonHandler(w http.ResponseWriter, r *http.Request)  {
	defer r.Body.Close()

	if r.Method == "POST" {
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}
		filename := fmt.Sprintf("./%d.txt", person.ID)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Printf("end!\n")
	} else  if r.Method == "GET" {
		fmt.Println("get!")

		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id)

		filename := fmt.Sprintf("./%d.txt", id)
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		person := Person{
			ID: id,
			Name: string(b),
		}

		t.Execute(w, person)

	}

}

func main()  {
	fmt.Println("start!")

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/persons", PersonHandler)

	http.ListenAndServe(":3777", nil)

}
