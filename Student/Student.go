package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Student struct {
	StudentID int    `json:"studentId"`
	FirstName string `json:"studentFirstName"`
	LastName  string `json:"studentLastName"`
	Email     string `json:"studentEmail"`
	Password  string `json:"studentPassword"`
}

var (
	db  *sql.DB
	err error
)

func dB() {
	db, err = sql.Open("mysql", "student:etistudentpwd@tcp(127.0.0.1:3306)/student_db")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}

func main() {
	dB()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/student", StudentSignUpHandler).Methods("POST")
	router.HandleFunc("/api/v1/student", StudentLoginHandler).Methods("GET")

	fmt.Println("Listening at port 5212")
	http.ListenAndServe(":5212",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Origin", "X-Api-Key", "X-Requested-With", "Content-Type", "Accept", "Authorization"}),
			handlers.AllowCredentials(),
		)(router))
}

// create student acc/signup
func StudentSignUpHandler(w http.ResponseWriter, r *http.Request) {
	var newStudent Student
	err := json.NewDecoder(r.Body).Decode(&newStudent)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// // Insert the new student into the database
	stmt, err := db.Prepare("INSERT INTO Student (StudentID, FirstName, LastName, Email, Password) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(newStudent.StudentID, newStudent.FirstName, newStudent.LastName, newStudent.Email, newStudent.Password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Account created successfully")
}

func StudentLoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("studentEmail")
	password := r.URL.Query().Get("studentPassword")

	if email == "" || password == "" {
		http.Error(w, "Email and Password parameters are required", http.StatusBadRequest)
		return
	}

	var acc Student
	err := db.QueryRow("SELECT StudentID, FirstName, LastName, Email, Password FROM Student WHERE Email = ? AND Password = ?", email, password).Scan(&acc.StudentID, &acc.FirstName, &acc.LastName, &acc.Email, &acc.Password)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid Email or Password", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Respond with user information
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}
