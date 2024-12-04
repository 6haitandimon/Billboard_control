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
	fs := http.FileServer(http.Dir("./ADS"))
	r.PathPrefix("/ADS/").Handler(http.StripPrefix("/ADS/", fs))

	app.RegisterRoutes(r)
	corsRouter := app.CORSMiddleware(r)

	//fs := http.FileServer(http.Dir("./ADS"))
	//http.Handle("/ADS/", http.StripPrefix("/ADS/", fs))

	log.Println("Server is running on port 6550")
	log.Fatal(http.ListenAndServe("10.160.67.17:6550", corsRouter))
}
