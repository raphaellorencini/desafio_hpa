package main

import (
    "fmt"
    "log"
	"net/http"
	"math"
)

func loop() string {
	x := 0.0001
	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}
	
    return fmt.Sprint("Code.education Rocks!")
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, loop()+"\n")
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8000", nil))
}