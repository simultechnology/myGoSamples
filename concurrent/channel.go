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
