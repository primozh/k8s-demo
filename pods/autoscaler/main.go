package main

import (
	"fmt"
	"math"
	"net/http"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	x := 0.0001
	for i := 1; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Done! Result = %f", x)
}

func main() {

	http.HandleFunc("/", handleIndex)

	http.ListenAndServe(":8080", nil)
}
