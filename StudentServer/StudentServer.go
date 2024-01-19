package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strconv"

	//"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	//"golang.org/x/text/date"
)

// struct for the student object
type Student struct {
	StudentID int    `json:"StudentID"`
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	PhoneNo   string `json:"Phone No"`
	Email     string `json:"Email"`
}

var (
	db  *sql.DB
	err error
)

// main server functions that handles all the microservices
func main() {
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/education")

	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/education/student/{studentID}", studentdata).Methods("GET", "DELETE", "POST", "PATCH", "PUT", "OPTIONS")
	router.HandleFunc("/api/v1/education/student", allstudent)

	fmt.Println("Listening at port 2479")
	log.Fatal(http.ListenAndServe(":2479", router))
}

// function that handles the student data, with methods to add, create, update and delete records.
func studentdata(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Update student data by student ID
	if r.Method == "POST" {
		println("POST Student")
		if body, err := io.ReadAll(r.Body); err == nil {
			var data Student
			fmt.Println(string(body))
			if err := json.Unmarshal(body, &data); err == nil {
				fmt.Println(params["studentID"])
				fmt.Println(reflect.TypeOf(params["studentID"]))
				if _, ok := isExist(params["studentID"]); !ok {
					fmt.Println(data)
					//courses[params["courseid"]] = data
					insertStudent(params["studentID"], data)

					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusConflict)
					fmt.Fprintf(w, "student ID exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "PUT" {
		// Update student data by student ID to replace the entire resource with a new representation
		println("PUT Student")
		if body, err := io.ReadAll(r.Body); err == nil {
			var data Student

			if err := json.Unmarshal(body, &data); err == nil {
				if _, ok := isExist(params["studentID"]); ok {
					fmt.Println(data)
					//courses[params["courseid"]] = data
					updateStudent(params["studentID"], data)
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "student ID does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "PATCH" {
		println("PATCH Student")
		// Patch student by student ID to apply partial updates to a resource

		if body, err := io.ReadAll(r.Body); err == nil {
			var data map[string]interface{}

			if err := json.Unmarshal(body, &data); err == nil {
				fmt.Println("Json")
				if orig, ok := isExist(params["studentID"]); ok {
					fmt.Println(data)

					for k, v := range data {
						switch k {
						case "First Name":
							orig.FirstName = v.(string)
						case "Last Name":
							orig.LastName = v.(string)
						case "Phone No":
							orig.PhoneNo = v.(string)
						case "Email":
							orig.Email = v.(string)
						}
					}
					//courses[params["courseid"]] = orig
					updateStudent(params["studentId"], orig)
					w.WriteHeader(http.StatusAccepted)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "student ID does not exist")
				}
			} else {
				fmt.Println(err)
			}
		}
	} else if r.Method == "DELETE" {
		if _, err := io.ReadAll(r.Body); err == nil {
			//var data Student
			println("DELETE Student func")
			//if err := json.Unmarshal(body, &data); err == nil {
			println(params["studentID"])
			if _, ok := isExist(params["studentID"]); ok {
				println("DELETE Student")
				// Delete student by ID

				fmt.Fprintf(w, params["studentID"]+" Deleted")
				//delete(courses, params["courseid"])
				delStudent(params["studentID"])
				// w.WriteHeader(http.StatusAccepted)

			} else {
				// input error
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "Invalid student ID")
			}
			//}
		} else {
			// input error
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Invalid student ID")
		}
	}
}

// Get all student information from database
func allstudent(w http.ResponseWriter, r *http.Request) {

	println("Query")
	db, _ := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/education")
	results, err := db.Query("select * from student")
	if err != nil {
		panic(err.Error())
	}

	var count int = 0

	//student := make(map[int]Student)
	var students map[int]Student = map[int]Student{}
	for results.Next() {
		var c Student

		_ = results.Scan(&c.StudentID, &c.FirstName, &c.LastName, &c.PhoneNo, &c.Email)

		count += 1
		students[count] = Student{StudentID: c.StudentID, FirstName: c.FirstName, LastName: c.LastName, PhoneNo: c.PhoneNo, Email: c.Email}
	}

	fmt.Println(count, " Records found")
	studentWrapper := struct {
		Students map[int]Student `json:"Student"`
	}{students}

	jsonBytes, err := json.Marshal(studentWrapper)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonBytes))

	json.NewEncoder(w).Encode(studentWrapper)

}

func getStudent() map[int]Student {
	results, err := db.Query("select * from student")
	if err != nil {
		panic(err.Error())
	}

	var students map[int]Student = map[int]Student{}

	for results.Next() {
		var c Student

		err = results.Scan(&c.StudentID, &c.FirstName, &c.LastName, &c.PhoneNo, &c.Email)
		if err != nil {
			panic(err.Error())
		}

		students[c.StudentID] = c
		println(students[c.StudentID].FirstName + " " + students[c.StudentID].LastName)
	}

	return students
}

// function checks if the student already exists - used in CREATE and UPDATE functions
func isExist(id string) (Student, bool) {
	var c Student
	println("isExist")

	//db, _ := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/education")

	var pid, _ = strconv.Atoi(id)
	println(pid)
	result := db.QueryRow("select * from student where studentID=?", pid)
	fmt.Println("QueryRow")
	err := result.Scan(&id, &c.FirstName, &c.LastName, &c.PhoneNo, &c.Email)
	//err := result.Scan(&c.StudentID)
	println("id:", c.StudentID)
	//println("First Name: ", c.FirstName)
	if err == sql.ErrNoRows {
		fmt.Println("Found!")
		return c, false

	}
	fmt.Println("Not Found!")
	return c, true
}

// function to delete a student from the server
func delStudent(id string) (int64, error) {
	fmt.Println("DeleteStudent func")
	var pid, _ = strconv.Atoi(id)
	fmt.Println("pid= ", pid)
	result, err := db.Exec("delete from student where studentID=?", pid)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// function to insert a student into the server database
func insertStudent(id string, c Student) {
	println("Insert Student")
	fmt.Println(" c.StudentID= ", c.StudentID)
	fmt.Println(" c.FirstName= ", c.FirstName)
	fmt.Println(" c.LastName= ", c.LastName)
	fmt.Println(" c.PhoneNo= ", c.PhoneNo)
	fmt.Println(" c.Email= ", c.Email)

	_, err := db.Exec("insert into student values(?,?,?,?,?)", c.StudentID, c.FirstName, c.LastName, c.PhoneNo, c.Email)
	if err != nil {
		panic(err.Error())
	}
}

// function to update a student record in the server database
func updateStudent(id string, c Student) {
	var pid, _ = strconv.Atoi(id)
	fmt.Println("Exec", pid)
	_, err := db.Exec("update student set studentID=?, firstName=?, lastName=?, phoneNo=?, email=? where studentID=?", pid, c.FirstName, c.LastName, c.PhoneNo, c.Email, pid)
	if err != nil {
		panic(err.Error())
	}
}

// function to send a student query to the server
func queryStudent(query string) (map[int]Student, bool) {
	//results, err := db.Query(curl http://localhost:2479/api/v1/education/student?q="select"+"*"+"from"+"student"+"where"+"firstName="+"\"Goh\""m student where lower(firstName) like lower(?) or lower(lastName) Like lower(?)", "%"+query+"%",  "%"+query+"%")
	// curl http://localhost:2479/api/v1/education/student?q="Select"+"*"+"from"+"student"
	// curl http://localhost:2479/api/v1/education/student?q="Select%20*%20from%20student"
	db, _ := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/education")
	results, err := db.Query(query)
	println("queryStudent routine: ", &results)

	if err != nil {
		panic(err.Error())
	}

	var students map[int]Student = map[int]Student{}

	for results.Next() {
		var c Student
		//var id string
		err = results.Scan(&c.StudentID, &c.FirstName, &c.LastName, &c.PhoneNo, &c.Email)
		if err != nil {
			panic(err.Error())
		}
		println("StudentID: " + strconv.Itoa(c.StudentID))
		students[c.StudentID] = c
	}

	if len(students) == 0 {
		return students, false
	}
	return students, true
}

// function to find eligible students that fit the query
func findEligibleStudent(name string) (map[int]Student, bool) {
	results, err := db.Query("select * from student where firstName like ?", name)
	if err != nil {
		panic(err.Error())
	}

	var student map[int]Student = map[int]Student{}

	for results.Next() {
		var c Student
		//var id string
		err = results.Scan(&c.StudentID, &c.FirstName, &c.LastName, &c.PhoneNo, &c.Email)
		if err != nil {
			panic(err.Error())
		}

		student[c.StudentID] = c
	}

	if len(student) == 0 {
		return student, false
	}
	return student, true
}
