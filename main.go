package main

import (
	"net/http"

	"github.com/simpleInternetUser/TaskManager/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Index)
	http.ListenAndServe(":8080", nil)
}
