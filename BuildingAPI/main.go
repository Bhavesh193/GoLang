package main

import (
	"encoding/json" // handles converting Go structs into JSON and back
	"fmt"           // printing
	"log"
	"math/rand"
	"net/http" // allow to to create http server
	"strconv"
	"time"

	"github.com/gorilla/mux" // third party router that helps you to handle URL path
)

// Encoding = Converting Go data into JSON (for sending to client).
// Decoding = Converting JSON into Go data (for reading from request).

type Course struct {
	CourseId     string  `json:"courseid"`
	CoursesName  string  `json:"coursename"`
	CoursesPrice int     `json:"courseprice"`
	Author       *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

// Middleware , helper - file
func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CoursesName == ""
}

func main() {
	fmt.Println("API - CRUD Operation")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "10", CoursesName: "ReactJS", CoursesPrice: 500, Author: &Author{FullName: "Bhavesh", Website: "bhavesh.com"}})
	courses = append(courses, Course{CourseId: "20", CoursesName: "Java", CoursesPrice: 700, Author: &Author{FullName: "Raju", Website: "raju.com"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Code</h1>"))
}

// w → The HTTP response writer (used to send data to the client).
// Encode(courses) → Converts the Go courses slice to JSON and sends it.

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loops through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send some Data")
		return
	}

	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("No Data inside JSON")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[:index+1]...)

			var course Course

			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			// courses = append(courses[:index], courses[:index+1]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[:index+1]...)
			break
		}
	}
}
