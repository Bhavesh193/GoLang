package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice string `json:"courseprice"`
	Author *Author `json:"author"`
}
type Author struct {
	FullName string `json:"fullname"`
	Website string `json:"autor"`
}

var courses []Course

func (c *Course) isEmpty() bool {
return c.CourseId == "" && c.CourseName ==""
}

func serveHome (w http.ResponseWriter , r *http.Request) {
	w.Write([]byte("<h1> Welcome to API for the crud </h1>"))
}

func getAllCourses( w http.ResponseWriter , r *http.Request){
	fmt.Println("Get All Courses")
	w.Header().Set("content-type" , "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourses( w http.ResponseWriter , r *http.Request){
	fmt.Println("Get One Courses")
	w.Header().Set("content-type" , "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

func createOneCourse( w http.ResponseWriter , r *http.Request){
	fmt.Println("Create One Course")
	w.Header().Set("content-type" , "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send Some Data")
		return
	}
	
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if(course.isEmpty()){
		json.NewEncoder(w).Encode("No Data inside Json")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses , course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse( w http.ResponseWriter , r *http.Request){
	fmt.Println("Create One Course")
	w.Header().Set("content-Type" , "application/json")

	params := mux.Vars(r) ;

	for index , course := range courses {
		if course.CourseId == params["id"]{
			courses = append(courses[:index],courses[:index + 1]... )

			var course  Course 

			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func main() {
	fmt.Println("Crud")
}
