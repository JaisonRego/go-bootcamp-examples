package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const HOST = "localhost:8000"

type course struct {
	Name     string `json:"CoarseName"`
	Price    int
	Platform string   `json:"Website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Making Web resuqests to a frontend Server")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	// PerformPostFormRequest()
	EncodeJson()
	// DecodeJson()
}

func getBaseUrl(path string) string {
	finalUrl := &url.URL{
		Scheme: "http",
		Host:   HOST,
		Path:   path,
	}
	return finalUrl.String()
}

func PerformGetRequest() {
	response, err := http.Get(getBaseUrl("get"))
	checkError(err)
	defer response.Body.Close()

	fmt.Println("Status Code is: ", response.StatusCode)
	fmt.Println("Content Length is: ", response.ContentLength)

	var responseData strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	checkError(err)
	byteCount, _ := responseData.Write(content)

	fmt.Println("Byte Count is: ", byteCount)
	fmt.Println(responseData.String())
}

func PerformPostJsonRequest() {
	// payload := strings.NewReader(`
	// 	{
	// 		"User": "Jaison",
	// 		"CoarseName": "Angular",
	// 		"Price": 10
	// 	}
	// `)
	response, err := http.Post(getBaseUrl("post"), "application/json", strings.NewReader(EncodeJson()))
	checkError(err)
	defer response.Body.Close()

	fmt.Println("Status Code is: ", response.StatusCode)
	fmt.Println("Content Length is: ", response.ContentLength)

	var responseData strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	checkError(err)

	byteCount, _ := responseData.Write(content)
	fmt.Println("Byte Count is: ", byteCount)
	fmt.Println(responseData.String())
}

func PerformPostFormRequest() {
	data := url.Values{}
	data.Add("firstname", "jaison")
	data.Add("lastname", "rego")
	data.Add("email", "jaison@co.in")

	response, err := http.PostForm(getBaseUrl("postform"), data)
	checkError(err)
	defer response.Body.Close()

	fmt.Println("Status Code is: ", response.StatusCode)
	fmt.Println("Content Length is: ", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)
	checkError(err)
	fmt.Println(string(content))
}

func EncodeJson() string {
	data := []course{
		{"React js Bootcamp", 199, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 599, "LearnCodeOnline.in", "abc123", nil},
	}

	jsonData, err := json.MarshalIndent(data, "", "\t")
	checkError(err)

	// var jsonDataFinal strings.Builder
	// jsonDataFinal.Write(jsonData)
	// fmt.Println(jsonDataFinal.String())

	fmt.Println(string(jsonData))
	return string(jsonData)
}

func DecodeJson() {
	// jsonDataFromWeb := []byte(`
	// 	{
	// 		"CoarseName": "React js Bootcamp",
	// 		"Price": 199,
	// 		"Website": "LearnCodeOnline.in",
	// 		"tags": ["web-dev","js"]
	// 	}
	// `)
	jsonDataFromWeb := []byte(EncodeJson())
	var coursedata []course
	if json.Valid(jsonDataFromWeb) {
		json.Unmarshal(jsonDataFromWeb, &coursedata)
		fmt.Printf("%+v\n", coursedata)
	} else {
		fmt.Println("Invalid Json Data")
	}

	// var custom []map[string]interface{}
	// json.Unmarshal(jsonDataFromWeb, &custom)
	// for _, data := range custom {
	// 	for key, value := range data {
	// 		fmt.Printf("The key is %v and value is %v\n", key, value)
	// 	}
	// }
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
