package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Tutor struct {
	TutorID   int    `json:"tutorId"`
	FirstName string `json:"tutorFirstName"`
	LastName  string `json:"tutorLastName"`
	Email     string `json:"tutorEmail"`
	Password  string `json:"tutorPassword"`
}

func main() {
outer:
	for {
		fmt.Println("===============================================")
		fmt.Println("Welcome!")
		fmt.Println("1. Create Tutor Account")
		fmt.Println("2. Login")
		fmt.Println("0. Exit")
		fmt.Print("Enter an option: ")

		var choice int
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			// user account creation
			fmt.Println("----Create Tutor Account----")
			createTutorAcc()
		case 2:
			// user login
			fmt.Println("----Login----")
			err := login()
			if err != nil {
				fmt.Println("Login failed:", err)
				return
			}
		case 0:
			break outer
		default:
			fmt.Println("Invalid option")
		}
		fmt.Scanln()
	}
}

// creates user account
func createTutorAcc() {
	var tutor Tutor
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	fmt.Print("Enter Firstname: ")
	fmt.Scanf("%v", &tutor.FirstName)
	reader.ReadString('\n')
	fmt.Print("Enter Lastname: ")
	fmt.Scanf("%v", &tutor.LastName)
	reader.ReadString('\n')
	fmt.Print("Enter Email: ")
	fmt.Scanf("%v", &tutor.Email)
	reader.ReadString('\n')
	fmt.Print("Enter Password: ")
	fmt.Scanf("%v", &tutor.Password)

	postBody, _ := json.Marshal(tutor)

	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodPost, "http://localhost:5211/api/v1/tutor", bytes.NewBuffer(postBody)); err == nil {
		if res, err := client.Do(req); err == nil {
			if res.StatusCode == 201 {
				fmt.Println("Account request sent. Please wait for admin approval.")
			} else {
				fmt.Println("Error creating user account")
			}
		} else {
			fmt.Println(2, err)
		}
	} else {
		fmt.Println(3, err)
	}

}

func login() error {
	var (
		email    string
		password string
	)
	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n')
	fmt.Print("Enter Email: ")
	fmt.Scanf("%v", &email)

	reader.ReadString('\n')
	fmt.Print("Enter Password: ")
	fmt.Scanf("%v", &password)

	// Perform login check
	client := &http.Client{}
	if req, err := http.NewRequest(http.MethodGet, "http://localhost:5211/api/v1/tutor?tutorEmail="+email+"&tutorPassword="+password, nil); err == nil {
		if res, err := client.Do(req); err == nil {
			defer res.Body.Close()

			if res.StatusCode == http.StatusOK {
				var tutor Tutor
				err := json.NewDecoder(res.Body).Decode(&tutor)
				if err == nil {
					fmt.Printf("Welcome back, %s!\n", tutor.Email)
					return nil
				} else {
					return fmt.Errorf("Error decoding response: %v", err)
				}
			} else {
				return fmt.Errorf("Inavlid Email or Password")
			}
		} else {
			return fmt.Errorf("Error making request: %v", err)
		}
	} else {
		return fmt.Errorf("Error creating request: %v", err)
	}
}
