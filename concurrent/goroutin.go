package main
import (
	"fmt"
	"net/http"
	"log"
	"sync"
)

func main()  {
	fmt.Println("start!")

	wait := new(sync.WaitGroup)
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
		wait.Add(1)
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			fmt.Println(url, res.Status)
			wait.Done()
		}(url)
	}
	wait.Wait()
}
