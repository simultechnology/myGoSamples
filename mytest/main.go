package main

import (
	"fmt"
	"github.com/guregu/kami"
	"golang.org/x/net/context"
	"net/http"
	"flag"
)

func hello(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", kami.Param(ctx, "name"))
}
func main() {
	flag.Set("bind", ":1234")
	kami.Get("/hello/:name", hello)
	kami.Serve()
}
