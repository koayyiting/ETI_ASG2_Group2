package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Tutor struct {
	TutorID   int    `json:"tutorId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

var (
	db  *sql.DB
	err error
)

func dB() {
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/carpooling_db")
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
	router.HandleFunc("/api/v1/tutor/{tutorID}", updateTutorProfileHandler).Methods("PUT")
	router.HandleFunc("/api/v1/tutor/{tutorID}", deleteTutorAccHandler).Methods("DELETE")

	fmt.Println("Listening at port 2211")
	log.Fatal(http.ListenAndServe(":2211", router))
}

func createTutorAccHandler(w http.ResponseWriter, r *http.Request) {
	var newTutor Tutor
	err := json.NewDecoder(r.Body).Decode(&newTutor)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// // Insert the new tutor into the database
	stmt, err := db.Prepare("INSERT INTO User (TutorID, FirstName, LastName, Email) VALUES (?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(newTutor.TutorID, newTutor.FirstName, newTutor.LastName, newTutor.Email)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Account created successfully")
}

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
	stmt, err := db.Prepare("UPDATE Tutor SET FirstName=?, LastName=?,  Email=? WHERE TutorID=?")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(updatedTutor.FirstName, updatedTutor.LastName, updatedTutor.Email, tutorID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	fmt.Fprintln(w, "Tutor profile updated successfully!")
}

// Add the deleteUserHandler function
func deleteTutorAccHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the request URL
	vars := mux.Vars(r)
	tutorID, err := strconv.Atoi(vars["tutorID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Delete the tutor from the database
	stmt, err := db.Prepare("DELETE FROM Tutor WHERE TutorID=?")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(tutorID)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Tutor Account deleted successfully!")
}
