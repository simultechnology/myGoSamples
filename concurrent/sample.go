package main
import (
	"fmt"
	"net/http"
	"log"
)

func main()  {
	fmt.Println("start!")

	urls := []string{
		"https://www.mapbox.com/",
		"http://52.69.128.185/", "http://google.co.jp",
		"http://52.69.128.185/", "http://google.co.jp",
		"http://52.69.128.185/", "http://google.co.jp",
		"http://52.69.128.185/", "http://google.co.jp",
		"http://52.69.128.185/", "http://google.co.jp",
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
