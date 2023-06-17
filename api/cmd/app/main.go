package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	baseUrl := os.Getenv("BASE_URL")
	port := os.Getenv("PORT")

	http.HandleFunc(baseUrl+"/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Hello, world!")
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
