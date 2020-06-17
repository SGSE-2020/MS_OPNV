package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	userpb "main/internal/proto"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var USERNAME = os.Getenv("POSTGRES_USER")
var PASSWORD = os.Getenv("POSTGRES_PASSWORD")
var DBNAME = os.Getenv("POSTGRES_DB")
var DB_HOST = os.Getenv("DB_HOST")
var GRPC_HOST = "ms-buergerbuero"
var API_PORT = "8080"
var GRPC_PORT = "50051"

var db *gorm.DB //database
var dbInitFlag bool = false

var grpc_client *grpc.ClientConn

type User struct {
	gorm.Model
	UId   string `json:"uid"`
	Token string `json:"token"`
}

type Area struct {
	gorm.Model
	Name     string `json:"name"`
	AreaType string `json:"areaType"`
	Price    int    `json:"price"`
}

type Ticket struct {
	gorm.Model
	AreaId       string `json:"areaId"`
	Qrcode       string `json:"qrCode"`
	Validitydate string `json:"validityDate"`
	TicketType   string `json:"ticketType"`
}

type TaxiCompany struct {
	gorm.Model
	Name       string `json:"name"`
	Adress     string `json:"adress"`
	Telenumber string `json:"telenumber"`
}

type TrafficInformation struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}

type VerifyUser struct {
	Token string
}

func init() {
	ConnectDB()
	if dbInitFlag {
		defer GetDB().Close()
	}
}

func testPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the API!")
	//	fmt.Println("Endpoint Hit: /api")
}

func validateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user VerifyUser
	json.Unmarshal(reqBody, &user)
	json.NewEncoder(w).Encode(user)

	w.Header().Set("Content-Type", "application/json")

	ConnectGRPC()
	client := userpb.NewUserServiceClient(grpc_client)
	ctx := context.Background()
	verifiedUser, err := client.VerifyUser(ctx, &userpb.UserToken{Token: user.Token})
	if err != nil {
		w.Write([]byte("{\"User ID\": \"Der gRPC Call VerifyUser hat nicht geklappt\"}"))
	} else {
		userData, err := client.GetUser(ctx, &userpb.UserId{Uid: verifiedUser.Uid})
		if err != nil {
			w.Write([]byte("{\"User ID\": \"Der gRPC Call GetUser hat nicht geklappt\"}"))
		} else {
			jsonData, err := json.Marshal(userData)
			if err != nil {
				log.Println(err)
			}
			w.Write([]byte(string(jsonData)))
		}
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

	// active routes
	myRouter.HandleFunc("/api", testPage)
	myRouter.HandleFunc("/users", getUsers)
	myRouter.HandleFunc("/user", validateUser).Methods("POST")

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

func ConnectGRPC() {
	conn, _ := grpc.Dial(
		GRPC_HOST+":"+GRPC_PORT, grpc.WithInsecure())
	grpc_client = conn
	// if err != nil {
	// 	fmt.Print("Keine Verbindung zum grpc-Server")
	// 	fmt.Print(err)
	// 	return false
	// } else {
	// 	fmt.Println("Ich gehe hier trotzdem rein")
	// 	grpc_client = conn
	// 	return true
	// }
}
