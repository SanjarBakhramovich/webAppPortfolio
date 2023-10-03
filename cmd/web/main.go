package main

import (
	"fmt"
	"myGoWebApp/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting server at port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
