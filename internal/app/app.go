package app

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/hello", HelloPagesHandler).Methods("GET")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/registration", RegistrationHandler).Methods("POST")
	r.HandleFunc("/ads/view", GetDeviceADSHandler).Methods("POST")
	// Защищенные маршруты пользователь
	userProtecte := r.PathPrefix("/user").Subrouter()
	userProtecte.Use(AuthMiddleware)
	userProtecte.HandleFunc("/device/list", UserDeviceListHandler).Methods("GET")
	userProtecte.HandleFunc("/device/add/check", GetFreeDevicesHandler).Methods("GET")
	userProtecte.HandleFunc("/device/add", AddUserDeviceHandler).Methods("POST")
	userProtecte.HandleFunc("/device/group", GetDevicesGroupHandler).Methods("GET")
	userProtecte.HandleFunc("/device/group/add", AddDevicesToGroupHandler).Methods("POST")
	userProtecte.HandleFunc("/device/group/delete", DeleteDevicesToGroupHandler).Methods("POST")
	userProtecte.HandleFunc("/device/group/setup", GetDeviceToGroupHandler).Methods("POST")
	userProtecte.HandleFunc("/device/group/create", DeviceGroupCreateHandler).Methods("POST")
	userProtecte.HandleFunc("/schedule/list", GetScheduleHandler).Methods("GET")
	userProtecte.HandleFunc("/schedule/setting", ScheduleSettingHandler).Methods("POST")
	userProtecte.HandleFunc("/schedule/setting/update", ScheduleSettingHandler).Methods("POST")
	userProtecte.HandleFunc("/schedule/setting/update/save", ScheduleSettingSaveHandler).Methods("POST")
	userProtecte.HandleFunc("/media", GetADSAllHandler).Methods("GET")
	userProtecte.HandleFunc("/statistic", GetStatisticHandler).Methods("POST")

	//adsProtecte := r.PathPrefix("/ads").Subrouter()
	//adsProtecte.Use(AuthMiddleware)

	// Защищенные маршруты админ
	adminProtected := r.PathPrefix("/admin").Subrouter()
	adminProtected.Use(AuthMiddleware)
	adminProtected.HandleFunc("/dashboard", AdminDashboardHandler).Methods("GET")
	adminProtected.HandleFunc("/users", GetUsersHandler).Methods("GET")
	adminProtected.HandleFunc("/user/delete", DeleteUsersHandler).Methods("POST")
	adminProtected.HandleFunc("/devices", GetDevicesHandler).Methods("GET")
	adminProtected.HandleFunc("/device/add", AddAdminDeviceHandler).Methods("POST")
	adminProtected.HandleFunc("/device/delete", DeleteDeviceHandler).Methods("POST")
	adminProtected.HandleFunc("/log/get", GetLogsHandler).Methods("POST")
	adminProtected.HandleFunc("/media", GetADSAllHandler).Methods("GET")
	adminProtected.HandleFunc("/media/delete", DeleteADSAllHandler).Methods("POST")
	//adminProtected.HandleFunc("/devices_update", UpdateDiveceHander).Methods("PATCH")
}
