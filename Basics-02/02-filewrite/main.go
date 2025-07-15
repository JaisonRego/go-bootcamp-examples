package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This will be inside a file , user is {{ .Name }} whose age is {{ .Age}} \n"

	file, _ := os.Create("./mytestfile.txt")

	len, _ := io.WriteString(file, content)

	fmt.Println(len)
	user := User{"jaison", 23}
	pathfile := "/Volumes/Development/App/Go-Lang/BootCamp/Basics-02/filewrite/mytestfile.txt"
	t, _ := template.New("mytestfile.txt").ParseFiles(pathfile)

	t.Execute(os.Stdout, user)

	defer file.Close()
	data, _ := ioutil.ReadFile("./mytestfile.txt")
	fmt.Println(string(data))
}
