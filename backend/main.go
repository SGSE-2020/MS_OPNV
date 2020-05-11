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

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password int    `json:"pw"`
}

func init() {

	// e := godotenv.Load("../env/.env") //Load .env file
	// if e != nil {
	// 	fmt.Print(e)
	// }

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&User{}) //Database migration
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func createNewUser(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"hello\": \"world\"}"))
	fmt.Println("user api Reched")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	json.NewEncoder(w).Encode(user)
	name := user.Name
	email := user.Email
	pw := user.Password
	GetDB().Create(&User{Password: int(pw), Name: string(name), Email: string(email)})

}

func handleRequests() {
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8000" //localhost
	}
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/user", createNewUser).Methods("POST")
	handler := cors.Default().Handler(myRouter)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func main() {
	handleRequests()
}

func GetDB() *gorm.DB {
	return db
}
