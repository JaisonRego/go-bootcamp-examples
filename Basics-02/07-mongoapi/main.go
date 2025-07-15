package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jaisonrego/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")

	fmt.Println("Server is Starting...")
	router := router.Router()
	log.Fatal(http.ListenAndServe(":4000", router))
	fmt.Println("Listening at port 4000...")
}
