package main

import (
	"log"
	"main/pkg/api"
	"net/http"
)

func main() {
	api.FillEndpoints()
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
