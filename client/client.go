package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	for {
		go func() {
			resp, err := http.Get("http://localhost:5959/api/test")
			if err != nil {
				panic(err)
			}

			defer resp.Body.Close()
			fmt.Println("Response status:", resp.Status)
		}()

		time.Sleep(2 * time.Second)
	}
}
