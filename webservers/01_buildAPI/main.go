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
	fmt.Println("Api for my backend")

	//making router
	r := mux.NewRouter()

	//seeding -> inserting data into database
	courses = append(courses, Course{CourseId: "2", CourseName: "Golang + microservices", CoursePrice: 899, Author: &Author{Fullname: "Sahil Sharma", Website: "Sahil.dev"}})
	courses = append(courses, Course{CourseId: "1", CourseName: "DevOps", CoursePrice: 599, Author: &Author{Fullname: "Sahil Sharma", Website: "Sahil.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN bootcamp", CoursePrice: 799, Author: &Author{Fullname: "Sahil Sharma", Website: "youtube"}})

	//routing
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(":4000", r))

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
	// return
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
	err := json.NewDecoder(r.Body).Decode(&course)

	if err != nil {
		json.NewEncoder(w).Encode("Invalid data")
		return
	}

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the JSON")
		return
	}

	//TODO: to check if course title is duplicate
	for _, existingCourse := range courses {
		if existingCourse.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course already exists, give different name")
			return
		}
	}

	//generate a unique id, convert to string
	//append new course into courses

	rand.NewSource(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100)) //convert rand int into string
	courses = append(courses, course)              //added

	json.NewEncoder(w).Encode(course) //response
	// return
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

// delete the one course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop, id , remove -index, index + 1

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("deleted the course with given Id..")
			break //breakes the looop
		}
	}
}
