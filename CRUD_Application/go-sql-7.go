package main

import (
	"net/http"
)

func main() {
	routers()
	http.ListenAndServe(":8005", Logger())
}