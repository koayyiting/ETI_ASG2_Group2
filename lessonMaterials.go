package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type LessonMaterial struct {
	TutorID int    `json:"TutorID"`
	Topic   string `json:"Topic"`
	Summary string `json:"Summary"`
	Created string `json:"Created on"`
}

var (
	db          *sql.DB
	err         error
	lm          LessonMaterial = LessonMaterial{}
	currentTime time.Time      = time.Now()
)

func main() {

	// [Edit] SQL Connection
	db, err = sql.Open("mysql", "etiLessonMaterial:eti2024@tcp(127.0.0.1:3306)/lessonmaterial_db")
	if err != nil {
		panic(err.Error())
	}

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/lessonmaterial/{materialid}", material).Methods("POST", "DELETE", "OPTIONS")
	router.HandleFunc("/lessonmaterial/all", allmaterials)
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

// REST API Functions
func allmaterials(w http.ResponseWriter, r *http.Request) {
	materialJSON := struct {
		Materials map[string]LessonMaterial `json: "Lesson Materials"`
	}{getMaterials()}

	json.NewEncoder(w).Encode(materialJSON)
}

func material(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("Lesson Material Function")

	//Adding New Materials
	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			fmt.Println(r.Method + " Method")
			if err := json.Unmarshal(body, &lm); err == nil {
				fmt.Println(string(body))

				if _, ok := isExist(params["materialid"]); !ok {
					fmt.Print("Lesson Material : ", lm)

					addMaterial(params["materialid"], lm)
					w.WriteHeader(http.StatusAccepted)
					fmt.Println(r.Method, " Response: ", r.Response)

				} else {
					w.WriteHeader(http.StatusConflict)
					fmt.Fprint(w, "Lesson Materials exists")
					fmt.Println(r.Method, " Response: ", r.Response)
				}
			} else {
				fmt.Print(err)
			}

		}
	} else if r.Method == "PUT" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			fmt.Println(r.Method + " Method")
			if err := json.Unmarshal(body, &lm); err == nil {
				fmt.Println("Object: ", lm)
				if _, ok := isExist(params["materialid"]); ok {

					editMaterial(params["materialid"], lm)
					fmt.Println(r.Method, " Response: ", r.Response)
					w.WriteHeader(http.StatusAccepted)
				} else {

					w.WriteHeader(http.StatusNotFound)
					fmt.Println(r.Method, " Response: ", r.Response)
				}
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println(err)
		}
	} else if lmaterial, ok := isExist(params["materialid"]); ok {
		if r.Method == "DELETE" {
			fmt.Println(r.Method + " Method")

			deleteMaterial(params["materialid"])

			fmt.Fprintf(w, params["materialid"]+" deleted")
		} else {
			json.NewEncoder(w).Encode(lmaterial)
		}
	} else {
		fmt.Fprintf(w, "Invalid Material ID")
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(r.Response)
	}
}

// [Edit] SQL Functions
func getMaterials() map[string]LessonMaterial {

	//Return Lesson Materials
	var materials map[string]LessonMaterial = map[string]LessonMaterial{}

	lessonMtrls, err := db.Query("SELECT * FROM LessonMaterials")
	if err != nil {
		panic(err.Error())
	}

	for lessonMtrls.Next() {
		var lm LessonMaterial
		var lmID string

		// [Edit] Lesson Material Values
		_ = lessonMtrls.Scan(&lmID, &lm.TutorID, &lm.Topic, &lm.Summary, &lm.Created)

		//Adding to Materials Map
		materials[lmID] = lm
	}

	return materials
}

func isExist(id string) (LessonMaterial, bool) {
	var lm LessonMaterial

	// [Edit] Lesson Material Values
	result := db.QueryRow("SELECT * FROM LessonMaterials WHERE ID=?", id)
	err := result.Scan(&id, &lm.TutorID, &lm.Topic, &lm.Summary)
	if err == sql.ErrNoRows {
		return lm, false
	}

	return lm, true
}
func addMaterial(id string, lm LessonMaterial) {
	// [Edit] Lesson Material Values
	_, err := db.Exec("INSERT INTO LessonMaterials VALUE (?,?,?,?,?)", id, lm.TutorID, lm.Topic, lm.Summary, currentTime)
	if err != nil {
		panic(err.Error())
	}
}

func editMaterial(id string, lm LessonMaterial) {
	// [Edit] Lesson Material Values
	_, err := db.Exec("UPDATE LessonMaterials SET TutorID=? Topic=? Summary=? WHERE ID=?", lm.TutorID, lm.Topic, lm.Summary, id)
	if err != nil {
		panic(err.Error)
	}
}

func deleteMaterial(id string) (int64, error) {
	// [Edit] Lesson Material Values
	result, err := db.Exec("DELETE FROM LessonMaterials WHERE ID=?", id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
