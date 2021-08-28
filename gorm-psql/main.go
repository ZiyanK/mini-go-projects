package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}
	fmt.Println(db)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Working")
}

func handleRequests() {
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
