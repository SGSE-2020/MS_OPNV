package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	account "main/internal/bank"
	user "main/internal/buergerbuero"
	parkplatz "main/internal/parkplatz"

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
var GRPC_HOST_BB = "ms-buergerbuero"
var GRPC_HOST_BANK = "ms-bank"
var GRPC_HOST_PP = "ms-parkplatz"
var API_PORT = "8080"
var GRPC_PORT = "50051"
var DEST_IBAN = "DE 23 1520 0000 1812 8252 34"

var db *gorm.DB //database
var dbInitFlag bool = false

var grpc_client *grpc.ClientConn

type User struct {
	gorm.Model
	UId string `json:"uId"`
}

type Area struct {
	gorm.Model
	AreaType string  `json:"areaType"`
	Price    float32 `json:"price"`
}

type Ticket struct {
	gorm.Model
	UId          string    `json:"uId"`
	AreaType     string    `json:"areaType"`
	Qrcode       string    `json:"qrCode"`
	Validitydate time.Time `json:"validityDate"`
	TicketType   int       `json:"ticketType"`
}

type BuyTicketRequest struct {
	gorm.Model
	UId        string `json:"uId"`
	AreaType   string `json:"areaType"`
	TicketType int    `json:"ticketType"`
}

func init() {
	ConnectDB()
	if dbInitFlag {
		defer GetDB().Close()
	}
}

func createDB(w http.ResponseWriter, r *http.Request) {
	if ConnectDB() {
		var resultArea Area

		if err := GetDB().Where("area_type = ?", "SB-Zone-1").First(&resultArea).Error; err != nil {
			fmt.Println(err)
			if err := GetDB().Create(&User{UId: "1234"}).Error; err != nil {
				w.Write([]byte("{\"Response\": \"User Konnte nicht erstellt werden\"}"))
				w.Write([]byte(fmt.Sprintf("{\"Error\": \"%s\"}", err)))
			}
			if err := GetDB().Create(&Area{AreaType: "SB-Zone-1", Price: 2.5}).Error; err != nil {
				w.Write([]byte("{\"Response\": \"User Konnte nicht erstellt werden\"}"))
				w.Write([]byte(fmt.Sprintf("{\"Error\": \"%s\"}", err)))
			}
			if err := GetDB().Create(&Area{AreaType: "SB-Zone-2", Price: 1.5}).Error; err != nil {
				w.Write([]byte("{\"Response\": \"User Konnte nicht erstellt werden\"}"))
				w.Write([]byte(fmt.Sprintf("{\"Error\": \"%s\"}", err)))
			}
			if err := GetDB().Create(&Area{AreaType: "SB-Zone-3", Price: 2.0}).Error; err != nil {
				w.Write([]byte("{\"Response\": \"User Konnte nicht erstellt werden\"}"))
				w.Write([]byte(fmt.Sprintf("{\"Error\": \"%s\"}", err)))
			}
			if err := GetDB().Create(&Area{AreaType: "B-Zone-1", Price: 1.5}).Error; err != nil {
				w.Write([]byte("{\"Response\": \"User Konnte nicht erstellt werden\"}"))
				w.Write([]byte(fmt.Sprintf("{\"Error\": \"%s\"}", err)))
			}
			if err := GetDB().Create(&Area{AreaType: "B-Zone-2", Price: 2.5}).Error; err != nil {
				w.Write([]byte("{\"Response\": \"User Konnte nicht erstellt werden\"}"))
				w.Write([]byte(fmt.Sprintf("{\"Error\": \"%s\"}", err)))
			}
			if err := GetDB().Create(&Area{AreaType: "B-Zone-3", Price: 2.0}).Error; err != nil {
				w.Write([]byte("{\"Response\": \"User Konnte nicht erstellt werden\"}"))
				w.Write([]byte(fmt.Sprintf("{\"Error\": \"%s\"}", err)))
			}
			if err := GetDB().Create(&Area{AreaType: "B-Zone-4", Price: 3.0}).Error; err != nil {
				w.Write([]byte("{\"Response\": \"User Konnte nicht erstellt werden\"}"))
				w.Write([]byte(fmt.Sprintf("{\"Error\": \"%s\"}", err)))
			} else {
				fmt.Fprintf(w, "Die Daten wurden erstellt!")
				w.Write([]byte("{\"Response\": \"User Konnte nicht erstellt werden\"}\""))
				w.Write([]byte("{\"Response\": \"Und noch ein Zweites\"}\""))
			}
		} else {
			fmt.Fprintf(w, "Die Daten existieren bereits!")
		}
		defer GetDB().Close()
	}
}

func updateParkingspace(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ConnectGRPC(GRPC_HOST_PP)
	client := parkplatz.NewParkplatzClient(grpc_client)
	ctx := context.Background()
	utils, err := client.Utilization(ctx, &parkplatz.UtilizationRequest{ServiceName: "Test"})
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		w.Write([]byte(fmt.Sprintf("[")))
		for {
			w.Write([]byte(fmt.Sprintf("{")))
			res, err := utils.Recv()
			if err == io.EOF {
				defer grpc_client.Close()
				w.Write([]byte(fmt.Sprintf("}]")))
				return
			}
			if err != nil {
				json.NewEncoder(w).Encode(err)
			}
			w.Write([]byte(fmt.Sprintf("\"DisplayName\": \"%s\",", res.GetDisplayName())))
			w.Write([]byte(fmt.Sprintf("\"TotalSpots\": %d,", res.GetTotalSpots())))
			w.Write([]byte(fmt.Sprintf("\"UtilizedSpots\": %d", res.GetUtilizedSpots())))
			w.Write([]byte(fmt.Sprintf("},")))
		}
	}
	defer grpc_client.Close()
}

func fiveParkspaces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ConnectGRPC(GRPC_HOST_PP)
	client := parkplatz.NewParkplatzClient(grpc_client)
	ctx := context.Background()
	utils, err := client.Utilization(ctx, &parkplatz.UtilizationRequest{ServiceName: "Test"})
	if err != nil {
		json.NewEncoder(w).Encode(err)
		w.Write([]byte(fmt.Sprintf("[")))
	} else {
		w.Write([]byte(fmt.Sprintf("[")))
		for i := 0; i < 5; i++ {
			w.Write([]byte(fmt.Sprintf("{")))
			res, err := utils.Recv()
			if err == io.EOF {
				defer grpc_client.Close()
				w.Write([]byte(fmt.Sprintf("}]")))
				return
			}
			if err != nil {
				json.NewEncoder(w).Encode(err)
			}
			w.Write([]byte(fmt.Sprintf("\"DisplayName\": \"%s\",", res.GetDisplayName())))
			w.Write([]byte(fmt.Sprintf("\"TotalSpots\": %d,", res.GetTotalSpots())))
			w.Write([]byte(fmt.Sprintf("\"UtilizedSpots\": %d", res.GetUtilizedSpots())))
			w.Write([]byte(fmt.Sprintf("},")))
		}
	}
	defer grpc_client.Close()
}

func validateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var temp_user user.UserToken

	json.Unmarshal(reqBody, &temp_user)

	w.Header().Set("Content-Type", "application/json")
	ConnectGRPC(GRPC_HOST_BB)
	client := user.NewUserServiceClient(grpc_client)
	ctx := context.Background()
	verifiedUser, err := client.VerifyUser(ctx, &user.UserToken{Token: temp_user.Token})
	if err != nil {
		fmt.Println("Der gRPC Call VerifyUser hat nicht geklappt")
	} else {
		userData, err := client.GetUser(ctx, &user.UserId{Uid: verifiedUser.Uid})
		if err != nil {
			fmt.Println("Der gRPC Call GetUser hat nicht geklappt")
		} else {
			// test if user exists
			err := json.NewEncoder(w).Encode(userData)
			if err != nil {
				fmt.Println(err)
			} else if ConnectDB() {
				var resultUser User
				if err := GetDB().Where("uid = ?", userData.Uid).First(&resultUser).Error; err != nil {
					GetDB().Create(&User{UId: string(userData.Uid)})
					fmt.Println("User wurde erstellt")
				} else {
					fmt.Println("Der User existiert bereits")
				}
				defer GetDB().Close()
			}
		}
	}
	defer grpc_client.Close()
}

func buyTicket(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var ticketReq BuyTicketRequest
	json.Unmarshal(reqBody, &ticketReq)
	json.NewEncoder(w).Encode(ticketReq)

	if ConnectDB() {
		var resultUser User

		if err := GetDB().Where("uid = ?", ticketReq.UId).First(&resultUser).Error; err != nil {
			w.Write([]byte("{\"Response\": \"false at Where Uid\"}"))
			fmt.Println(err)
		} else {
			var now = time.Now()
			var valDate time.Time
			var temp_area Area
			var bill_sum float32
			if err := GetDB().Where("area_type = ?", ticketReq.AreaType).First(&temp_area).Error; err != nil {
				fmt.Println("false at Where Area")
				fmt.Println(err)
			} else {
				if ticketReq.TicketType == 0 {
					valDate = now.AddDate(0, 0, 1)
					bill_sum = temp_area.Price
				} else {
					valDate = now.AddDate(0, 1, 0)
					var bill_sum_temp = temp_area.Price * 0.2
					bill_sum = bill_sum_temp * 30
				}
			}
			if err := GetDB().Create(
				&Ticket{
					UId:          string(ticketReq.UId),
					AreaType:     ticketReq.AreaType,
					Qrcode:       "DummyQRCode",
					Validitydate: valDate,
					TicketType:   ticketReq.TicketType}).Error; err != nil {
				fmt.Println("false at Create Ticket")
			} else {
				ConnectGRPC(GRPC_HOST_BANK)
				client := account.NewAccountServiceClient(grpc_client)
				ctx := context.Background()
				acc, err := client.GetIban(ctx, &account.UserId{
					UserId: ticketReq.UId,
				})
				if err != nil {
					fmt.Println("Der gRPC Call GetIban hat nicht geklappt")
				} else {
					var temp_amount string
					temp_amount = fmt.Sprintf("%f", bill_sum)
					message, err := client.Transfer(ctx, &account.Transfer{
						UserId:    ticketReq.UId,
						Iban:      acc.Iban,
						Purpose:   "Ticket gekauft",
						DestIban:  DEST_IBAN,
						Amount:    temp_amount,
						StartDate: "",
						Repeat:    "",
					})
					if err != nil {
						fmt.Println("Der gRPC Call Transfer hat nicht geklappt")
						fmt.Println(message)
					} else {
						w.Write([]byte(fmt.Sprintf("{\"bill\": %f}", bill_sum)))
					}
				}
			}
		}
		defer GetDB().Close()
	} else {
		w.Write([]byte("{\"Response\": \"false\"}"))
	}
	defer grpc_client.Close()
}

func getUser(w http.ResponseWriter, r *http.Request) {
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
	myRouter.HandleFunc("/db", createDB)
	myRouter.HandleFunc("/user", validateUser).Methods("POST")
	myRouter.HandleFunc("/buy", buyTicket).Methods("POST")
	myRouter.HandleFunc("/parkspace", updateParkingspace)
	myRouter.HandleFunc("/fiveparkspaces", updateParkingspace)

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
			GetDB().Debug().AutoMigrate(&User{}, &Area{}, &Ticket{}) //Database migratio
			dbInitFlag = true
		}
		return true
	}
}

func ConnectGRPC(host string) {
	conn, _ := grpc.Dial(
		host+":"+GRPC_PORT, grpc.WithInsecure())
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
