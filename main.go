package main

import (
	"fmt"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// begin web page
	var hostPlatform = os.Getenv("HOST_PLATFORM")
	//var hostPlatform = "Azure Container Instance"

	var htmlHeader = "<!DOCTYPE html><html><h2>Simple web app</h2>"
	fmt.Fprintf(w, htmlHeader)
	fmt.Fprintf(w, "<body><p>Running on: %s</p></body></html>", hostPlatform)

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
