package main
import (
	"fmt"
	"net/http"
	"log"
)

func main()  {
	fmt.Println("start!")

	urls := []string{
		"http://192.168.33.99/avs/", "http://localhost:60011/testleaflet.html",
		"http://localhost:60011/testleaflet.html", "http://192.168.33.99/avs/",
		"http://localhost:60011/testleaflet.html", "http://192.168.33.99/avs/",
		"http://localhost:60011/testleaflet.html", "http://localhost:60011/testleaflet.html",
		"http://192.168.33.99/avs/", "http://localhost:60011/testleaflet.html",
		"http://192.168.33.99/avs/", "http://localhost:60011/testleaflet.html",
		"http://localhost:60011/testleaflet.html", "http://192.168.33.99/avs/",
		"http://localhost:60011/testleaflet.html", "http://192.168.33.99/avs/",
		"http://localhost:60011/testleaflet.html", "http://192.168.33.99/avs/",
		"http://localhost:60011/testleaflet.html", "http://192.168.33.99/avs/",
	}
	for _, url := range urls {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		fmt.Println(url, res.Status)
	}
}
