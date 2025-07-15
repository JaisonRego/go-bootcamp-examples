package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseID    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

var cources []Course

func (c Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("API - APP Starts Here")
	r := mux.NewRouter()

	cources = append(cources, Course{CourseID: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{FullName: "Jaison", Website: "lco.dev"}})
	cources = append(cources, Course{CourseID: "3", CourseName: "Angular", CoursePrice: 499, Author: &Author{FullName: "Alex", Website: "lco.dev"}})
	cources = append(cources, Course{CourseID: "5", CourseName: "Golang", CoursePrice: 199, Author: &Author{FullName: "Adam", Website: "lco.dev"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	r.HandleFunc("/deleteAllCourse", deleteAllCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API based application</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cources)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, course := range cources {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found with current ID")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adding a new course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
	}

	for _, value := range cources {
		if value.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Coarse Already Exist")
			return
		}
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	cources = append(cources, course)
	json.NewEncoder(w).Encode("Course sucessfully added")
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one value")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range cources {
		if course.CourseID == params["id"] {
			cources = append(cources[:index], cources[index+1:]...)

			var updatedCourse Course
			updatedCourse.CourseID = params["id"]
			json.NewDecoder(r.Body).Decode(&updatedCourse)

			cources = append(cources, updatedCourse)
			json.NewEncoder(w).Encode("Sucessfully updated the course")
			return
		}
	}
	json.NewEncoder(w).Encode("Course ID not found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Record")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, course := range cources {
		if course.CourseID == params["id"] {
			cources = append(cources[:index], cources[index+1:]...)
			json.NewEncoder(w).Encode("Course Sucessfully deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("Coarse ID not found")
}

func deleteAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Record")
	w.Header().Set("Content-Type", "application/json")
	cources = cources[:0]
	json.NewEncoder(w).Encode("All Courses are deleted")
}
