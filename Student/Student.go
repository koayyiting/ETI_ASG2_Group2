package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	//"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//"golang.org/x/text/date"
)

// struct for the student object
type Student struct {
	StudentID        int    `json:"StudentID"`
	StudentFirstName string `json:"StudentFirstName"`
	StudentLastName  string `json:"StudentLastName"`
	PhoneNo          string `json:"PhoneNo"`
	StudentEmail     string `json:"StudentEmail"`
	StudentPassword  string `json:"StudentPassword"`
}

var (
	db  *sql.DB
	err error
)

// main server functions that handles all the microservices
func main() {
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/studentDB")

	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/student", getStudent).Methods("GET")
	router.HandleFunc("/api/v1/student/{StudentID}", updateStudent).Methods("PATCH")
	router.HandleFunc("/api/v1/student", createStudent).Methods("POST")
	// router.HandleFunc("/api/v1/student/", deleteStudent).Methods("DELETE")
	fmt.Println("Listening at port 3306")
	http.ListenAndServe(":3306",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Origin", "X-Api-Key", "X-Requested-With", "Content-Type", "Accept", "Authorization"}),
			handlers.AllowCredentials(),
		)(router))
}

// function that creates student records and add them to the database.
func createStudent(w http.ResponseWriter, r *http.Request) {
	var newStudent Student
	err := json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// // Insert record into database
	stmt, err := db.Prepare("INSERT INTO Student (StudentID, StudentFirstName, StudentLastName, PhoneNo, StudentEmail, StudentPassword) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Server Error 500", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(newStudent.StudentID, newStudent.StudentFirstName, newStudent.StudentLastName, newStudent.PhoneNo, newStudent.StudentEmail, newStudent.StudentPassword)
	if err != nil {
		http.Error(w, "Server Error 500", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Account Created Successfully")
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("studentEmail")
	password := r.URL.Query().Get("studentPassword")

	if email == "" || password == "" {
		http.Error(w, "Server Error 400 - Bad Request", http.StatusBadRequest)
		return
	}

	var acc Student
	err := db.QueryRow("SELECT StudentID, StudentFirstName, StudentLastName, Phone, StudentEmail, StudentPassword FROM Student WHERE StudentEmail = ? AND Password = ?", email, password).Scan(&acc.StudentID, &acc.StudentFirstName, &acc.StudentLastName, &acc.PhoneNo, &acc.StudentEmail, &acc.StudentPassword)
	if err == sql.ErrNoRows {
		http.Error(w, "Incorrect Email or Password", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Server Error 500", http.StatusInternalServerError)
		return
	}

	// Respond with user information
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	// Parse the student ID from the request URL
	vars := mux.Vars(r)
	StudentID, err := strconv.Atoi(vars["StudentID"])
	if err != nil {
		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}

	var updStudent Student
	err = json.NewDecoder(r.Body).Decode(&updStudent)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Update student information in database
	stmt, err := db.Prepare("UPDATE Student SET StudentFirstName=?, StudentLastName=?, PhoneNo=?, StudentEmail=?, StudentPassword=? WHERE StudentID=?")
	if err != nil {
		http.Error(w, "Server Error 500", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(updStudent.StudentFirstName, updStudent.StudentLastName, updStudent.PhoneNo, updStudent.StudentEmail, updStudent.StudentPassword, StudentID)
	if err != nil {
		http.Error(w, "Server Error 500", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, "Profile successfully updated.")
}

// function to delete a student from the server
//func deleteStudent(id string) (int64, error) {
//fmt.Println("Delete Student Account")
//var pid, _ = strconv.Atoi(id)
//fmt.Println("pid= ", pid)
//result, err := db.Exec("delete from Student where StudentID=?", pid)
//if err != nil {
// 		return 0, err
// 	}
// 	return result.RowsAffected()
// }

// // function checks if the student already exists - used in CREATE and UPDATE functions
// func isExist(id string) (Student, bool) {
// 	var c Student
// 	println("isExist")

// 	//db, _ := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/studentDB")

// 	var pid, _ = strconv.Atoi(id)
// 	println(pid)
// 	result := db.QueryRow("select * from student where StudentID=?", pid)
// 	fmt.Println("QueryRow")
// 	err := result.Scan(&id, &c.StudentFirstName, &c.StudentLastName, &c.PhoneNo, &c.StudentEmail)
// 	//err := result.Scan(&c.StudentID)
// 	println("id:", c.StudentID)
// 	//println("First Name: ", c.StudentFirstName)
// 	if err == sql.ErrNoRows {
// 		fmt.Println("Found!")
// 		return c, false

// 	}
// 	fmt.Println("Not Found!")
// 	return c, true
// }
