package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for courses - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` //Author will be the type pointer
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// mocking databases here
//fake DB

var courses []Course

// middleware / helper func -> they are seperate file
// will process when both the filed courseid and course name is present
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

//controllers -> used to be in own seperate file

// servr Home route
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to our shop!!</h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All the courses")
	//setting headers
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) //In Encode what ever we provide is treated as json value and throwed back to w
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r) //all the request will give data
	fmt.Println(params)

	//loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given Id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what if: entire body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about the data -> {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the JSON")
		return
	}

	//generate a unique id, convert to string
	//append new course into courses

	rand.NewSource(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100)) //convert rand int into string
	courses = append(courses, course)              //added

	json.NewEncoder(w).Encode(course) //response
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//first -grab id from req
	params := mux.Vars(r)

	// loop through the value, id and remove , and add again with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	id, ok := params["id"]
	if !ok || id == "" {
		json.NewEncoder(w).Encode("Please provide the id")
		return
	}
}
