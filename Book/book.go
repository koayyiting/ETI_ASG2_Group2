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

type Booking struct {
	StudentName  string `json:"Student Name"`
	StudentEmail string `json:"Student Email"`
	LessonID     string `json:"Lesson ID"` //link to schedule details
	LessonName   string `json:"LessonName"`
	Location     string `json:"Location"`
}

var (
	db  *sql.DB
	err error
)

func main() {
	router := mux.NewRouter()
	router.Use(corsMiddleware)
	router.HandleFunc("/api/v1/book", newBook).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/book/{bookingID}", existBooking).Methods("GET", "DELETE", "POST", "OPTIONS")
	fmt.Println("Listening at port 1765")
	log.Fatal(http.ListenAndServe(":1765", router))
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
	db, err = sql.Open("mysql", "booking_system:ETI_Group2_Booking@tcp(127.0.0.1:3306)/ETI_Booking")

	if err != nil {
		panic(err.Error())
	}
}

func newBook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		retrieveBookings("") // need use session to get the email of student
	case http.MethodPost: //create booking
		if body, err := io.ReadAll(r.Body); err == nil {
			var newBooking Booking
			fmt.Println(string(body))
			if err := json.Unmarshal(body, &newBooking); err == nil {
				fmt.Println(newBooking)
				if createErr := createBooking(newBooking); createErr == nil {
					w.WriteHeader(http.StatusAccepted) // 202
				} else {
					w.WriteHeader(http.StatusConflict)
					fmt.Println("Error creating booking: ", createErr)
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

// create new booking
func createBooking(b Booking) error {
	fmt.Println("In createBooking function")
	openDB()
	defer db.Close()
	_, err := db.Exec("insert into Booking (StudentName, StudentEmail, LessonID) values(?,?,?)", b.StudentName, b.StudentEmail, b.LessonID)
	if err != nil {
		return err
	}
	return nil
}

// retrieve all booking done by student
// mb should do lesson name
func retrieveBookings(email string) error {
	fmt.Println("In retrieveBookings function")
	openDB()
	defer db.Close()
	_, err := db.Exec("select LessonID from Booking where StudentEmail = ?", email)
	if err != nil {
		return err
	}
	return nil
}

func existBooking(w http.ResponseWriter, r *http.Request) {
	bookingID := mux.Vars(r)["bookingID"]

	switch r.Method {
	case http.MethodPut:
		if body, err := io.ReadAll(r.Body); err == nil {
			var updatedBooking Booking
			if err := json.Unmarshal(body, &updatedBooking); err == nil {
				json.NewDecoder(r.Body).Decode(&updatedBooking)
				if err := updateBooking(updatedBooking); err == nil {
					w.WriteHeader(http.StatusAccepted) //202
				} else {
					w.WriteHeader(http.StatusConflict)
				}

			} else {
				fmt.Println(err)
			}
		}
	case http.MethodDelete:
		if status := deleteBooking(bookingID); status {
			w.WriteHeader(http.StatusAccepted) //202
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	}
}

// update Booking details
// theres nothing much to update for now..
func updateBooking(updatedBooking Booking) error {
	fmt.Println("In updateBooking func")
	openDB()
	defer db.Close() //will run at the end of the block of the code
	_, err := db.Exec("update Booking set StudentName=?, StudentEmail=?, where ID=?;", updatedBooking.StudentEmail, updatedBooking.StudentEmail)
	if err != nil {
		return err
	}
	return nil
}

// delete booking
func deleteBooking(id string) bool {
	fmt.Println("in deleteBooking function")
	openDB()
	defer db.Close() //will run at the end of the block of the code
	_, err := db.Exec("delete from Booking where ID = ?", id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// for curl on cmd

// var testing map[string]Booking = map[string]Booking{
// 	"Test_ID": {StudentName: "koay yi ting", StudentEmail: "koayyiting@gmail.com", LessonID: "test_lesson_id"},
// }

// func book(w http.ResponseWriter, r *http.Request) {
// 	bookingID := mux.Vars(r)["bookingID"]
// 	switch r.Method {
// 	case http.MethodGet:
// 		chosenBooking, _ := testing[bookingID]
// 		json.NewEncoder(w).Encode(chosenBooking)
// 	case http.MethodPost: //create booking
// 		var newBooking Booking
// 		err := json.NewDecoder(r.Body).Decode(&newBooking)
// 		if err != nil {
// 			fmt.Fprintf(w, "Invalid json format\n")
// 			return
// 		}
// 		testing[bookingID] = newBooking
// 		w.WriteHeader(http.StatusAccepted)
// 	}
// }
