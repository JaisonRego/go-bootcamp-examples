package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const url1 string = "https://lco.dev"
const url2 string = "https://lco.dev:3000/learn?coursename=reactjs&payment=ghj145jsh"

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Making a http request")

	responce, err := http.Get(url1)
	checkError(err)

	fmt.Printf("The responce is of type %T\n", responce)
	defer responce.Body.Close()

	databytes, err := io.ReadAll(responce.Body)
	checkError(err)

	fmt.Println(string(databytes))

	result, _ := url.Parse(url2)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Println(qparams)

	partsofUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=jaison",
	}
	fmt.Println(partsofUrl.String())
}
