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
		"https://www.mapbox.com/",
		"http://52.69.128.185/", "http://google.co.jp",
		"http://52.69.128.185/", "http://google.co.jp",
		"http://52.69.128.185/", "http://google.co.jp",
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
