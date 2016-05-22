package main
import (
	"fmt"
	"net/http"
	"log"
	"runtime"
	"time"
)

func main()  {
	fmt.Println("start!")

	startTime := time.Now().UnixNano() / int64(time.Millisecond)

	fmt.Println(startTime)

	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	urls := []string{
		"https://www.mapbox.com/",
		"http://52.69.128.185/", "http://google.co.jp",
		"http://52.69.128.185/", "http://google.co.jp",
		"http://52.69.128.185/", "http://google.co.jp",
	}

	statusChan := make(chan string)
	for _, url := range urls {
		go func(url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			statusChan <- res.Status
			fmt.Println(url, res.Status)
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Printf("i : %d\n", i)
		fmt.Println(<-statusChan)
	}

	endTime := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println(endTime)

	fmt.Printf("duration : %d mill sec", endTime - startTime)
}
