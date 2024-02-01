package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Booking struct {
	BookingID  int `json:"booking_id"`
	StudentID  int `json:"student_id"`
	ScheduleID int `json:"schedule_id"`
}

var (
	db  *sql.DB
	err error
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/book", newBook).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/getBookings", getBooking).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/oneBooking/{bookingID}", oneBooking).Methods("GET", "DELETE", "POST", "OPTIONS")
	fmt.Println("Listening at port 1765")
	http.ListenAndServe(":1765",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Origin", "X-Api-Key", "X-Requested-With", "Content-Type", "Accept", "Authorization"}),
			handlers.AllowCredentials(),
		)(router))
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
		retrieveBookings() // need use session to get the email of student
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
	_, err := db.Exec("INSERT into Booking (StudentID, ScheduleID) values(?,?)", b.StudentID, b.ScheduleID)
	if err != nil {
		return err
	}
	return nil
}

func getBooking(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if bookings, err := retrieveBookings(); err == nil {
			w.WriteHeader(http.StatusAccepted) //202
			if bookingsJSON, err := json.Marshal(bookings); err == nil {
				w.Write(bookingsJSON)
			} else {
				fmt.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	}
}

// retrieve all booking done by student
func retrieveBookings() ([]Booking, error) {
	fmt.Println("In retrieveBookings function")

	openDB()
	rows, err := db.Query("SELECT * FROM Booking")
	// rows, err := db.Query("SELECT * FROM Schedule WHERE TutorID = ?", tid)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var bookings []Booking

	for rows.Next() {
		var booking Booking
		if err := rows.Scan(&booking.BookingID, &booking.StudentID, &booking.ScheduleID); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}
	return bookings, nil
}

func oneBooking(w http.ResponseWriter, r *http.Request) {
	bookingID := mux.Vars(r)["bookingID"]

	switch r.Method {
	case http.MethodDelete:
		if status := deleteBooking(bookingID); status {
			w.WriteHeader(http.StatusAccepted) //202
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	}
}

// delete booking
func deleteBooking(id string) bool {
	fmt.Println("in deleteBooking function")
	openDB()
	defer db.Close() //will run at the end of the block of the code
	_, err := db.Exec("DELETE FROM Booking WHERE BookingID = ?", id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
