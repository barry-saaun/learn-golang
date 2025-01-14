package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const PORT = ":5001"

func Greet(writer io.Writer, name string) {
	// fmt.FPrintf("Hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetingHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Barry")
}

func main() {
	log.Fatal(http.ListenAndServe(PORT, http.HandlerFunc(MyGreetingHandler)))
}
