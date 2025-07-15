package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to golang modules!!")
	greeting()

	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeting() {
	fmt.Println("Welcome User!!!")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	web, _ := ioutil.ReadFile("./webpage.html")
	// w.Write([]byte("<h1>Welcome to my webpage</h1>"))
	w.Write(web)
}
