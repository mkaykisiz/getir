package main

import (
	"fmt"
	"getir/app"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	appModule := new(app.App)
	appModule.Initialize()

	http.HandleFunc("/", appModule.Handler)
	http.HandleFunc("/records", appModule.RecordListHandler)
	http.HandleFunc("/in-memory", appModule.InMemoryHandler)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Println("Server Listening on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
