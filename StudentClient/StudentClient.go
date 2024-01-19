package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Struct for Student
type Student struct {
	StudentID int    `json:"StudentID"`
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	PhoneNo   string `json:"Phone No"`
	Email     string `json:"Email"`
}

type Students struct {
	Students map[string]Student `json:"Student"`
}

// menu display for client service
func main() {

	var keyin int
	var quitprog = 1

	scanner := bufio.NewScanner(os.Stdin)

	for quitprog == 1 {
		mainmenu()
		// fmt.Scan(&keyin)
		scanner.Scan()
		keyin, _ = strconv.Atoi(scanner.Text())

		switch keyin {
		case 1: // Student
			studentmenu()
		case 9: // Return
			//  Quit
			quitprog = 0
		default:
			fmt.Println("### Invalid Input ###")
		}

	}
	os.Exit(0)
}

// function to list all students in the database
func listAllStudent() {
	//client := &http.Client{}
	resp, err := http.Get("http://localhost:2479/api/v1/education/student")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//var res Students
	var students Students

	err1 := json.Unmarshal(body, &students)
	if err != nil {
		fmt.Println(err1)
	}

	for k, v := range students.Students {
		fmt.Println("(", k, ") First Name: ", v.FirstName)
		fmt.Println("       Last Name: ", v.LastName)
		fmt.Println("       Phone: ", v.PhoneNo)
		fmt.Println("       Email: ", v.Email)
		fmt.Println()
	}

}

// function to create a new student
func createStudent() {
	var student Student
	var studentID string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the Student ID to be created: ")
	scanner.Scan()
	studentID = scanner.Text()
	student.StudentID, _ = strconv.Atoi(studentID)

	fmt.Print("First name: ")
	scanner.Scan()
	student.FirstName = scanner.Text()
	fmt.Print("Last Name: ")
	scanner.Scan()
	student.LastName = scanner.Text()
	fmt.Print("Phone no: ")
	scanner.Scan()
	student.PhoneNo = scanner.Text()
	fmt.Print("Email: ")
	scanner.Scan()
	student.Email = scanner.Text()

	postBody, _ := json.Marshal(student)
	resBody := bytes.NewBuffer(postBody)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:2479/api/v1/education/student/"+studentID, resBody); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Student", studentID, "created successfully")
			} else if res.StatusCode == 409 {
				fmt.Println("Error - Student", studentID, "exists")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

// function to update a student record
func updateStudent() {
	var student Student
	var studentID string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the Student ID to be created: ")
	scanner.Scan()
	studentID = scanner.Text()
	student.StudentID, _ = strconv.Atoi(studentID)
	fmt.Print("First name: ")
	scanner.Scan()
	student.FirstName = scanner.Text()
	fmt.Print("Last Name: ")
	scanner.Scan()
	student.LastName = scanner.Text()
	fmt.Print("Phone no: ")
	scanner.Scan()
	student.PhoneNo = scanner.Text()
	fmt.Print("Email: ")
	scanner.Scan()
	student.Email = scanner.Text()

	postBody, _ := json.Marshal(student)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPut, "http://localhost:2479/api/v1/education/student/"+studentID, bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 202 {
				fmt.Println("Student ", studentID, "updated successfully")
			} else if res.StatusCode == 404 {
				fmt.Println("Error - Student ", studentID, "does not exist")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

// function to delete a student record
func deleteStudent() {
	var studentID string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter the ID of the Student to be deleted: ")
	scanner.Scan()
	studentID = scanner.Text()

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodDelete, "http://localhost:2479/api/v1/education/student/"+studentID, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 200 {
				fmt.Println("Student", studentID, "deleted successfully")
			} else if res.StatusCode == 404 {
				fmt.Println("Error - student", studentID, "does not exist")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}
}

// main display for the client service
func mainmenu() {
	println("==========")
	println("Education Support System")
	println("  1.  Student")
	println("  9.  Quit")
	print("Enter an option: ")

}

// menu display for the student service
func studentmenu() {
	var keyin int
	scanner := bufio.NewScanner(os.Stdin)
loop:
	for {
		println("=================================")
		println("Student Support System")
		println("======== Student ==============")
		println("  1.  List Student Profiles")
		println("  2.  Create New Student")
		println("  3.  Update Student Profile")
		println("  4.  Delete Student Account")
		println("  9.  Back to Main Menu")
		print("Enter an option: ")

		scanner.Scan()
		keyin, _ = strconv.Atoi(scanner.Text())
		switch keyin {
		case 1:
			listAllStudent()
		case 2:
			createStudent()
		case 3:
			updateStudent()
		case 4:
			deleteStudent()
		case 9:
			//  Quit
			break loop
		default:
			fmt.Println("### Invalid Input ###")
		}
	}
}
