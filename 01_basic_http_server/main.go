package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_, _ = w.Write([]byte("Hello, from Go net/http server!"))
}

func main() {

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server Running on port 8080")

	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err)
}
