package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Schedule struct {
	ScheduleID     int       `json:"schedule_id"`
	ScheduleID_str string    `json:"schedule_id_str"`
	TutorID        int       `json:"tutor_id"`
	LessonID       int       `json:"lesson_id"`
	LessonID_str   int       `json:"lesson_id_str"`
	LessonName     string    `json:"lesson_name"`
	Location       string    `json:"location"`
	StartTime_str  string    `json:"start_time_str"`
	StartTime      time.Time `json:"start_time"`
	EndTime_str    string    `json:"end_time_str"`
	EndTime        time.Time `json:"end_time"`
}

var (
	db  *sql.DB
	err error
)

func main() {
	// openDB()
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/schedule", newSchedule).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/v1/getSchedules/{tid}", getSchedules).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/v1/getAllSchedules", getAllSchedules).Methods("GET")
	router.HandleFunc("/api/v1/oneSchedule/{sid}", oneSchedule).Methods("GET", "PUT", "DELETE", "OPTIONS")
	fmt.Println("Listening at port 1000")
	http.ListenAndServe(":1000",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Origin", "X-Api-Key", "X-Requested-With", "Content-Type", "Accept", "Authorization"}),
			handlers.AllowCredentials(),
		)(router))
}

func openDB() {
	db, err = sql.Open("mysql", "schedule_system:ETI_Group2_Schedule@tcp(127.0.0.1:3306)/ETI_Schedule")

	if err != nil {
		panic(err.Error())
	}
}

func getSchedules(w http.ResponseWriter, r *http.Request) {
	tid := mux.Vars(r)["tid"]
	switch r.Method {
	case http.MethodGet:
		if schedules, err := retrieveSchedules(tid); err == nil {
			w.WriteHeader(http.StatusAccepted) //202
			if schedulesJSON, err := json.Marshal(schedules); err == nil {
				w.Write(schedulesJSON)
			} else {
				fmt.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	}
}

func retrieveSchedules(tid string) ([]Schedule, error) {
	// func retrieveSchedules(tid int) ([]Schedule, error) {
	fmt.Println("In retrieveSchedules function")
	openDB()
	defer db.Close()
	rows, err := db.Query("SELECT ScheduleID, TutorID, LessonID, LessonName, Location, StartTime, EndTime FROM Schedule WHERE TutorID = ?", tid)
	// rows, err := db.Query("SELECT * FROM Schedule WHERE TutorID = ?", tid)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var schedules []Schedule
	var startTimeStr string
	var endTimeStr string

	for rows.Next() {
		var schedule Schedule
		if err := rows.Scan(&schedule.ScheduleID, &schedule.TutorID, &schedule.LessonID, &schedule.LessonName, &schedule.Location, &startTimeStr, &endTimeStr); err != nil {
			return nil, err
		}
		sgt, _ := time.LoadLocation("Asia/Singapore")
		schedule.StartTime, _ = time.ParseInLocation("2006-01-02 15:04:05", startTimeStr, sgt)
		schedule.EndTime, _ = time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, sgt)
		fmt.Println(schedule)
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func getAllSchedules(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if schedules, err := retrieveAllSchedules(); err == nil {
			w.WriteHeader(http.StatusAccepted) //202
			if schedulesJSON, err := json.Marshal(schedules); err == nil {
				w.Write(schedulesJSON)
			} else {
				fmt.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	}
}

func retrieveAllSchedules() ([]Schedule, error) {
	// func retrieveSchedules(tid int) ([]Schedule, error) {
	fmt.Println("In retrieveSchedules function")
	openDB()
	defer db.Close()
	rows, err := db.Query("SELECT ScheduleID, TutorID, LessonID, LessonName, Location, StartTime, EndTime FROM Schedule")
	// rows, err := db.Query("SELECT * FROM Schedule WHERE TutorID = ?", tid)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	var schedules []Schedule
	var startTimeStr string
	var endTimeStr string

	for rows.Next() {
		var schedule Schedule
		if err := rows.Scan(&schedule.ScheduleID, &schedule.TutorID, &schedule.LessonID, &schedule.LessonName, &schedule.Location, &startTimeStr, &endTimeStr); err != nil {
			return nil, err
		}
		sgt, _ := time.LoadLocation("Asia/Singapore")
		schedule.StartTime, _ = time.ParseInLocation("2006-01-02 15:04:05", startTimeStr, sgt)
		schedule.EndTime, _ = time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, sgt)
		fmt.Println(schedule)
		schedules = append(schedules, schedule)
	}
	return schedules, nil
}

func newSchedule(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In newSchedule function")

	switch r.Method {
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
	_, err := db.Exec("insert into Schedule (TutorID, LessonID, LessonName, Location, StartTime, EndTime) values(?,?,?,?,?,?)", s.TutorID, s.LessonID, &s.LessonName, &s.Location, s.StartTime_str, s.EndTime_str)
	if err != nil {
		return err
	}
	return nil
}

func oneSchedule(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in oneSchedule")
	id := mux.Vars(r)["sid"]
	switch r.Method {
	case http.MethodGet:
		if schedule, err := retrieveOneSchedule(id); err == nil {
			w.WriteHeader(http.StatusAccepted) //202
			if scheduleJSON, err := json.Marshal(schedule); err == nil {
				w.Write(scheduleJSON)
			} else {
				fmt.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	case http.MethodPut: //update status
		if body, err := io.ReadAll(r.Body); err == nil {
			var schedule Schedule
			if err := json.Unmarshal(body, &schedule); err == nil {
				json.NewDecoder(r.Body).Decode(&schedule)
				if err := updateSchedule(schedule); err == nil {
					fmt.Println(schedule)
					w.WriteHeader(http.StatusAccepted) //202
					scheduleJSON, _ := json.Marshal(schedule)
					w.Write(scheduleJSON)
				} else {
					w.WriteHeader(http.StatusConflict)
				}
			} else {
				fmt.Println(err)
			}
		}

	case http.MethodDelete: //cancel trip
		if err := deleteSchedule(id); err == nil { //param = userid
			fmt.Println("deleted schedule")
			w.WriteHeader(http.StatusAccepted) //202
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	}
}

func retrieveOneSchedule(scheduleID string) (Schedule, error) {
	fmt.Println("In retrieveOneSchedule function")
	openDB()
	defer db.Close()

	var schedule Schedule
	var startTimeStr string
	var endTimeStr string

	err := db.QueryRow("SELECT ScheduleID, TutorID, LessonID, LessonName, Location, StartTime, EndTime FROM Schedule WHERE ScheduleID = ?", scheduleID).Scan(&schedule.ScheduleID, &schedule.TutorID, &schedule.LessonID, &schedule.LessonName, &schedule.Location, &startTimeStr, &endTimeStr)
	if err != nil {
		return Schedule{}, err
	}

	sgt, _ := time.LoadLocation("Asia/Singapore")
	schedule.StartTime, _ = time.ParseInLocation("2006-01-02 15:04:05", startTimeStr, sgt)
	schedule.EndTime, _ = time.ParseInLocation("2006-01-02 15:04:05", endTimeStr, sgt)
	fmt.Println(schedule)

	return schedule, nil
}

func updateSchedule(schedule Schedule) error {
	fmt.Println("update schedule query")
	openDB()
	defer db.Close() //will run at the end of the block of the code
	fmt.Println(schedule.ScheduleID_str)
	_, err := db.Exec("UPDATE Schedule SET StartTime=?, EndTime=?, Location=?, LessonName=?, LessonID=? WHERE ScheduleID=?", schedule.StartTime_str, schedule.EndTime_str, schedule.Location, schedule.LessonName, schedule.LessonID_str, schedule.ScheduleID_str)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func deleteSchedule(id string) error {
	fmt.Println("in delete schedule function")
	openDB()
	defer db.Close() //will run at the end of the block of the code
	fmt.Println(id)
	_, err := db.Exec("DELETE FROM Schedule WHERE ScheduleID = ?", id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
