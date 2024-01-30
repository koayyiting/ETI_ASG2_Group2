package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Tutor struct {
	TutorID   int    `json:"tutorId"`
	Username  string `json:"tutorUsername"`
	Password  string `json:"tutorPassword"`
	Title     string `json:"tutorTitle"`
	FirstName string `json:"tutorFirstName"`
	LastName  string `json:"tutorLastName"`
}

var (
	db  *sql.DB
	err error
)

func dB() {
	db, err = sql.Open("mysql", "tutor:etitutorpwd@tcp(127.0.0.1:3306)/tutor_db")
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
	router.HandleFunc("/api/v1/tutor", createTutorAccHandler).Methods("POST")
	router.HandleFunc("/api/v1/tutor", GetTutorAccHandler).Methods("GET")
	router.HandleFunc("/api/v1/tutor/{tutorID}", updateTutorProfileHandler).Methods("PUT")

	fmt.Println("Listening at port 5211")
	http.ListenAndServe(":5211",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Origin", "X-Api-Key", "X-Requested-With", "Content-Type", "Accept", "Authorization"}),
			handlers.AllowCredentials(),
		)(router))
}

// create tutor acc/signup
func createTutorAccHandler(w http.ResponseWriter, r *http.Request) {
	var newTutor Tutor
	err := json.NewDecoder(r.Body).Decode(&newTutor)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// // Insert the new tutor into the database
	stmt, err := db.Prepare("INSERT INTO Tutor (TutorID, Username, Password, Title, FirstName, LastName) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(newTutor.TutorID, newTutor.Username, newTutor.Password, newTutor.Title, newTutor.FirstName, newTutor.LastName)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Account created successfully")
}

func GetTutorAccHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("tutorUsername")
	password := r.URL.Query().Get("tutorPassword")

	if username == "" || password == "" {
		http.Error(w, "Username and Password parameters are required", http.StatusBadRequest)
		return
	}

	var acc Tutor
	err := db.QueryRow("SELECT TutorID, Username, Password, Title, FirstName, LastName FROM Tutor WHERE Username = ? AND Password = ?", username, password).Scan(&acc.TutorID, &acc.Username, &acc.Password, &acc.Title, &acc.FirstName, &acc.LastName)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid Username or Password", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Respond with user information
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}

// update tutor profile
func updateTutorProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the tutor ID from the request URL
	vars := mux.Vars(r)
	tutorID, err := strconv.Atoi(vars["tutorID"])
	if err != nil {
		http.Error(w, "Invalid tutor ID", http.StatusBadRequest)
		return
	}

	var updatedTutor Tutor
	err = json.NewDecoder(r.Body).Decode(&updatedTutor)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Update the tutor's information in the database
	stmt, err := db.Prepare("UPDATE Tutor SET Username=?, Password=?, Title=?, FirstName=?, LastName=? WHERE TutorID=?")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(updatedTutor.Username, updatedTutor.Password, updatedTutor.Title, updatedTutor.FirstName, updatedTutor.LastName, tutorID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, "Tutor profile updated successfully!")
}
