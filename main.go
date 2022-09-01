package main

import (
	"fmt"
	router "hello/25-MongoApi/Router"
	"log"
	"net/http"
)

func main() {
	fmt.Println("MongoDB API")
	fmt.Println("Server is getting started")
	r :=  router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")
}