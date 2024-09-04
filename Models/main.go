package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Courses struct {
	CourseID    string  `json:"courseid" `
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author" `
}
type Author struct {
	AuthorName string `json:"authorname"`
	Website    string `json:"website"`
}

var courses []Courses

func getCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}
func getCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range courses {
		if item.CourseID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return

		}
	}
}
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range courses {
		if item.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(courses)
			return

		}

	}
}
func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course Courses
	_ = json.NewDecoder(r.Body).Decode(&course)
	course.CourseID = strconv.Itoa(rand.Intn(1000000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
}
func main() {
	r := mux.NewRouter()
	courses = append(courses, Courses{CourseID: "1", CourseName: "Golang", CoursePrice: 200, Author: &Author{AuthorName: "Bhargavi", Website: "https://8080/golang.com"}})
	courses = append(courses, Courses{CourseID: "2", CourseName: "Java", CoursePrice: 400, Author: &Author{AuthorName: "steve", Website: "https://8080/java.com"}})

	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/courses", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	fmt.Println("Starting Server at 8080 port")
	log.Fatal(http.ListenAndServe(":8080", r))
}
