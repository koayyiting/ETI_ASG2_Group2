package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
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
	router.HandleFunc("/lessonmaterial/material/{materialid}", material).Methods("POST", "DELETE", "OPTIONS")
	router.HandleFunc("/lessonmaterial/tutor/{tutorid}", tutormaterial).Methods("GET")
	router.HandleFunc("/material/{id}", specificmaterial).Methods("GET")
	router.HandleFunc("/lessonmaterial/all", allmaterials)
	fmt.Println("Listening at port 4088")
	http.ListenAndServe(":4088",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Origin", "X-Api-Key", "X-Requested-With", "Content-Type", "Accept", "Authorization"}),
			handlers.AllowCredentials(),
		)(router))
}

func material(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println("Lesson Material Function")

	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			fmt.Println(r.Method + " Method")
			if err := json.Unmarshal(body, &lm); err == nil {
				fmt.Println(string(body))

				if _, ok := isExist(params["materialid"]); !ok {
					fmt.Print("Lesson Material : ", lm)

					addMaterial(params["materialid"], lm)
					w.WriteHeader(http.StatusAccepted)
					fmt.Println(r.Method, " Response: ", r.Response.Status)

				} else {
					w.WriteHeader(http.StatusConflict)
					fmt.Fprint(w, "Lesson Materials exists")
					fmt.Println(r.Method, " Response: ", r.Response.Status)
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
					fmt.Println(r.Method, " Response: ", r.Response.Status)
					w.WriteHeader(http.StatusAccepted)
				} else {

					w.WriteHeader(http.StatusNotFound)
					fmt.Println(r.Method, " Response: ", r.Response.Status)
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

func specificmaterial(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	materialJSON := struct {
		LessonMaterials map[string]LessonMaterial `json:"Material"`
	}{getMaterial(params["id"])}

	json.NewEncoder(w).Encode(materialJSON)
}

// REST API Functions
func allmaterials(w http.ResponseWriter, r *http.Request) {

	materialJSON := struct {
		LessonMaterials map[string]LessonMaterial `json:"Materials"`
	}{getMaterials()}

	json.NewEncoder(w).Encode(materialJSON)
}

func tutormaterial(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var tutorid = params["tutorid"]

	materialJSON := struct {
		LessonMaterials map[string]LessonMaterial `json:"Material"`
	}{getTutorMaterial(tutorid)}

	json.NewEncoder(w).Encode(materialJSON)
}

// [Edit] SQL Functions
func getMaterial(id string) map[string]LessonMaterial {

	//Return Lesson Materials
	var onematerial map[string]LessonMaterial = map[string]LessonMaterial{}

	lessonMtrls, err := db.Query("SELECT * FROM LessonMaterials WHERE LMID=?", id)
	if err != nil {
		panic(err.Error())
	}

	for lessonMtrls.Next() {
		var lm LessonMaterial
		var lmID string

		// [Edit] Lesson Material Values
		_ = lessonMtrls.Scan(&lmID, &lm.TutorID, &lm.Topic, &lm.Summary, &lm.Created)

		//Adding to Materials Map
		onematerial[lmID] = lm
	}

	return onematerial
}

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

func getTutorMaterial(id string) map[string]LessonMaterial {

	//Return Lesson Materials
	var materials map[string]LessonMaterial = map[string]LessonMaterial{}

	lessonMtrls, err := db.Query("SELECT * FROM LessonMaterials WHERE TutorID=?", id)
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
	result := db.QueryRow("SELECT * FROM LessonMaterials WHERE LMID=?", id)
	err := result.Scan(&id, &lm.TutorID, &lm.Topic, &lm.Summary)
	if err == sql.ErrNoRows {
		return lm, false
	}

	return lm, true
}
func addMaterial(id string, lm LessonMaterial) {
	// [Edit] Lesson Material Values
	_, err := db.Exec("INSERT INTO LessonMaterials VALUE (?,?,?,?,?)", id, lm.TutorID, lm.Topic, lm.Summary, lm.Created)
	if err != nil {
		panic(err.Error())
	}
}

func editMaterial(id string, lm LessonMaterial) {
	// [Edit] Lesson Material Values
	_, err := db.Exec("UPDATE LessonMaterials SET TutorID=? Topic=? Summary=? WHERE LMID=?", lm.TutorID, lm.Topic, lm.Summary, id)
	if err != nil {
		panic(err.Error)
	}
}

func deleteMaterial(id string) (int64, error) {
	// [Edit] Lesson Material Values
	result, err := db.Exec("DELETE FROM LessonMaterials WHERE LMID=?", id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
