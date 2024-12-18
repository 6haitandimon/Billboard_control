package main

import (
	"Billboard/internal/app"
	"Billboard/pkg/database"
	"fmt"
	"log"
	"net/http"
	"os"

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
	ServerHostAddress := os.Getenv("HOST_ADDR")
	ServerPort := os.Getenv("HOST_PORT")

	log.Println("Server is running on", ServerHostAddress, ServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", ServerHostAddress, ServerPort), corsRouter))
}
