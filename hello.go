package main

import (
	"fmt"
	"log"
	"net/http"
)

func SomethingHandler(w http.ResponseWriter, r *http.Request) {
	ipAddress := r.RemoteAddr

	fmt.Fprintf(w, "Hello dickHead")
	log.Printf("This is the IPaddress %s", ipAddress)
}

func main() {
	port := ":8080"
	fmt.Printf("The server started at port %s", port)
	http.HandleFunc("/last", SomethingHandler)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Print("The server was not started")
	}
	fmt.Printf("The server started at port %s\n", port)
}
