package main

import (
	"fmt"
	"github.com/ColaWsl/fresh-go/something"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe(":8080", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! 2")
	fmt.Println("access")
}

func TestWget() {
	something.Wget()
}
