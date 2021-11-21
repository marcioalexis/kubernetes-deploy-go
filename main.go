package main

import (
	"fmt"
	"net/http"
	"log"
)

const webContent = "marcio-devops.tk:v99"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
