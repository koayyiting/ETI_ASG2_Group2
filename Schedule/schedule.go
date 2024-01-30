package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Schedule struct {
	StudentName  string `json:"Student Name"`
	StudentEmail string `json:"Student Email"`
	LessonID     string `json:"Lesson ID"` //link to schedule details
	// add lesson venue
}

var (
	db  *sql.DB
	err error
)

func main() {
	// openDB()
	router := mux.NewRouter()
	router.Use(corsMiddleware)
	router.HandleFunc("/api/v1/schedule", newSchedule).Methods("GET", "DELETE", "POST", "PATCH", "PUT", "OPTIONS")
	fmt.Println("Listening at port 1000")
	log.Fatal(http.ListenAndServe(":1000", router))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Api-Key, X-Requested-With, Content-Type, Accept, Authorization")
		next.ServeHTTP(w, r)
	})
}

func openDB() {
	db, err = sql.Open("mysql", "schedule_system:ETI_Group2_Schedule@tcp(127.0.0.1:3306)/ETI_Schedule")

	if err != nil {
		panic(err.Error())
	}
}

func newSchedule(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		retrieveSchedules("") // need use session to get the email of student
	case http.MethodPost: //create schedule
		if body, err := io.ReadAll(r.Body); err == nil {
			var newSchedule Schedule
			fmt.Println(string(body))
			if err := json.Unmarshal(body, &newSchedule); err == nil {
				fmt.Println(newSchedule)
				if createErr := createSchedule(newSchedule); createErr == nil {
					w.WriteHeader(http.StatusAccepted) // 202
				} else {
					w.WriteHeader(http.StatusConflict)
					fmt.Println("Error creating schedule: ", createErr)
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("Error unmarshaling JSON: ", err)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error reading request body: ", err)
		}
	}
}

func createSchedule(s Schedule) error {
	fmt.Println("In createSchedule function")
	openDB()
	defer db.Close()
	_, err := db.Exec("insert into Schedule (StudentName, StudentEmail, LessonID) values(?,?,?)", s.StudentName, s.StudentEmail, s.LessonID)
	if err != nil {
		return err
	}
	return nil
}

// retrieve all Schedule done by student
// mb should do lesson name
func retrieveSchedules(email string) error {
	fmt.Println("In retrieveSchedules function")
	openDB()
	defer db.Close()
	_, err := db.Exec("select scheduleID from Schedule where StudentEmail = ?", email)
	if err != nil {
		return err
	}
	return nil
}
