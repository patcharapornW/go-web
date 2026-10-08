package main

import (
	"fmt"
	"net/http"
)

func hellogoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go!")
}

func hiGoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, Go! eiei")
}

func main() {
	http.HandleFunc("/hellogo", hellogoHandler)
	http.HandleFunc("/higo", hiGoHandler)

	fmt.Printf("Server started at port 3030...")

	//start server
	if err := http.ListenAndServe(":3030", nil); err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
