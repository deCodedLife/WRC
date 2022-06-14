package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/info", HelloWorld).Methods("GET")

	http.ListenAndServe(":8821", router)
}
