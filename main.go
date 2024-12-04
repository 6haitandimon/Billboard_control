package main

import (
	"Billboard/internal/app"
	"Billboard/pkg/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	database.InitDB()

	r := mux.NewRouter()
	app.RegisterRoutes(r)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
