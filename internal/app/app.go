package app

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/registration", RegistrationHandler).Methods("POST")
	// Защищенные маршруты пользователь
	userProtecte := r.PathPrefix("/user").Subrouter()
	userProtecte.Use(AuthMiddleware)
	//userProtecte.

	// Защищенные маршруты админ
	adminProtected := r.PathPrefix("/admin").Subrouter()
	adminProtected.Use(AuthMiddleware)
	adminProtected.HandleFunc("/dashboard", AdminDashboardHandler).Methods("GET")
	adminProtected.HandleFunc("/users", GetUsersHandler).Methods("GET")
	adminProtected.HandleFunc("/user/delete", DeleteUsersHandler).Methods("POST")
	adminProtected.HandleFunc("/devices", GetDevicesHandler).Methods("GET")
	adminProtected.HandleFunc("/device/add", AddDeviceHandler).Methods("POST")
	adminProtected.HandleFunc("/device/delete", DeleteDeviceHandler).Methods("POST")
	adminProtected.HandleFunc("/log/get", GetLogsHandler).Methods("POST")
	adminProtected.HandleFunc("/media", GetADSAllHandler).Methods("GET")
	adminProtected.HandleFunc("/media/delete", DeleteADSAllHandler).Methods("POST")
	//adminProtected.HandleFunc("/devices_update", UpdateDiveceHander).Methods("PATCH")
}
