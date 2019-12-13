package main

import (
	"fmt"
	"log"
	"net/http"
)

var numberOfRequests int = 0

func apiTest(w http.ResponseWriter, req *http.Request) {
	numberOfRequests++
	fmt.Println(numberOfRequests)
}

func main() {
	port := 5959
	fmt.Printf("Server starts to listen in port %d\n", port)

	http.HandleFunc("/api/test", apiTest)
	log.Fatal(http.ListenAndServe(":5959", nil))
}
