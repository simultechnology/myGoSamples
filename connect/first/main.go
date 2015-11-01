package main
import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("start!!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		query := r.URL.RawQuery
		fmt.Fprint(w, query)
	})
	http.ListenAndServe(":3000", nil)
}
