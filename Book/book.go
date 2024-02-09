package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

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
	router.HandleFunc("/api/v1/book/{scheduleID}", newBook).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/getBookings/{sid}", getBooking).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/oneBooking/{bookingID}", oneBooking).Methods("DELETE", "OPTIONS")
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
	scheduleID := mux.Vars(r)["scheduleID"]
	switch r.Method {
	case http.MethodPost: //create booking
		if body, err := io.ReadAll(r.Body); err == nil {
			var newBooking Booking
			fmt.Println(string(body))
			if err := json.Unmarshal(body, &newBooking); err == nil {
				fmt.Println(newBooking)
				if notSameLesson := existingBooking(scheduleID); notSameLesson {
					if createErr := createBooking(newBooking); createErr == nil {
						w.WriteHeader(http.StatusAccepted) // 202
					} else {
						w.WriteHeader(http.StatusConflict)
						fmt.Println("Error creating booking: ", createErr)
					}
				} else {
					w.WriteHeader(http.StatusUnprocessableEntity)
					fmt.Println("You have already Booked this Schedule")
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

// check if booked before
func existingBooking(sid string) bool {
	fmt.Println("In createBooking function")
	openDB()
	defer db.Close()
	var booking Booking
	rowBooking := db.QueryRow("SELECT * FROM Booking where ScheduleID=?", sid)
	errBook := rowBooking.Scan(&booking.BookingID, &booking.StudentID, &booking.ScheduleID)
	if errBook != sql.ErrNoRows {
		fmt.Println("Booked Same Lesson")
		return false
	}
	return true
}

func getBooking(w http.ResponseWriter, r *http.Request) {
	studentId_str := mux.Vars(r)["sid"]
	studentId, _ := strconv.Atoi(studentId_str)
	switch r.Method {
	case http.MethodGet:
		if bookings, err := retrieveBookings(studentId); err == nil {
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
func retrieveBookings(studentId int) ([]Booking, error) {
	fmt.Println("In retrieveBookings function")

	openDB()
	rows, err := db.Query("SELECT * FROM Booking WHERE StudentID = ?", studentId)
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
