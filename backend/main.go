package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

var db *gorm.DB //database
var dbInitFlag bool = false
var USERNAME = os.Getenv("POSTGRES_USER")
var PASSWORD = os.Getenv("POSTGRES_PASSWORD")
var DBNAME = os.Getenv("POSTGRES_DB")

//var DBHOST = "192.168.99.102" //dev
var DB_HOST = os.Getenv("DB_HOST")
var API_PORT = "8080"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password int    `json:"pw"`
}

func init() {
	ConnectDB()
	if dbInitFlag {
		defer GetDB().Close()
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: /")
}

func testPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the API!")
	fmt.Println("Endpoint Hit: /api")

}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	json.NewEncoder(w).Encode(user)

	w.Header().Set("Content-Type", "application/json")
	if ConnectDB() {
		GetDB().Create(&User{Password: int(user.Password), Name: string(user.Name), Email: string(user.Email)})
		w.Write([]byte("{\"Response\": \"User wurde erstellt\"}"))
		defer GetDB().Close()
	} else {
		w.Write([]byte("{\"Response\": \"Keine Verbindung zur Datenbank\"}"))
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if ConnectDB() {
		var resultUsers []User
		GetDB().Find(&resultUsers)
		defer GetDB().Close()
		arr, err := json.Marshal(resultUsers)
		if err != nil {
			w.Write([]byte("{\"Response\": \"User konnte nicht Decodiert werden\"}"))
		} else {
			w.Write(arr)
		}
	} else {
		w.Write([]byte("{\"Response\": \"Keine Verbindung zur Datenbank\"}"))
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api", testPage)
	myRouter.HandleFunc("/users", getUsers)
	myRouter.HandleFunc("/user", createNewUser).Methods("POST")
	handler := cors.Default().Handler(myRouter)
	log.Fatal(http.ListenAndServe(":"+API_PORT, handler))
}

func main() {
	handleRequests()
}

// Database functions
func GetDB() *gorm.DB {
	return db
}

func ConnectDB() bool {
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", DB_HOST, USERNAME, DBNAME, PASSWORD) //Build connection string
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
		fmt.Print("Keine Verbindung zur Datenbank")
		return false
	} else {
		db = conn
		if !dbInitFlag {
			GetDB().Debug().AutoMigrate(&User{}) //Database migration
			dbInitFlag = true
		}
		return true
	}
}
