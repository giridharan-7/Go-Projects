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

type Student struct{
	ID string `json:"id"`
	Year string `json:"year"`
	Name *Name `json:"name"`
}
type Name struct{
	FirstName string `json:firstname`
	LastName string `json:lastname`
}

var students []Student

func getStudents(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func deleteStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range students{
		if item.ID == params["id"]{
			students = append(students[:index], students[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(students)
}

func getStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _,item := range students{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
}

func createStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var student Student
	_  = json.NewDecoder(r.Body).Decode(&student)
	student.ID = strconv.Itoa(rand.Intn(100000000))
	students = append(students, student)
	json.NewEncoder(w).Encode(student)
}

func updateStudent(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range students {
		if item.ID == params["id"]{
			students = append(students[:index],students[index+1:]...)
			var student Student
			_ = json.NewDecoder(r.Body).Decode(&student)
			student.ID = params["id"]
			students = append(students, student)
			json.NewEncoder(w).Encode(student)
		}
	}
}

func main(){
	r:= mux.NewRouter()

	students = append(students, Student{ID: "1", Year: "first", Name: &Name{FirstName: "Giridharan", LastName: "Pasupathi"}})
	students = append(students, Student{ID: "2", Year: "Second", Name: &Name{FirstName: "Madhav", LastName: "Reddy"}})

	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/students/{id}",getStudent).Methods("GET")
	r.HandleFunc("/students", createStudent).Methods("POST")
	r.HandleFunc("/students/{id}", updateStudent).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteStudent).Methods("DELETE")

	fmt.Printf("Starting the server at the port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}
