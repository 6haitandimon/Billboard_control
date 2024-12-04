package app

import (
	"Billboard/internal/repositories"
	"Billboard/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := services.FetchAllUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	//user, _ := repositories.GetUser(credentials.Username)

	token, err, ID := services.Authenticate(credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	//fmt.Println(strconv.Itoa(ID))

	response := map[string]string{"token": token, "RoleID": strconv.Itoa(ID)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var registration struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&registration); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	check, err := repositories.CheckUser(registration.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !check {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	err = services.Registration(registration.Username, registration.Password, 1)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Println("User is registration was successful")

}

func GetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	devices, err := services.FetchAllDevices()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(devices)
}

func AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(UserIDKey).(int)
	w.Write([]byte("Welcome, Admin! Your ID: " + string(rune(userID))))
}

func AddDeviceHandler(w http.ResponseWriter, r *http.Request) {
	ID := r.Context().Value(UserIDKey).(int)

	device, err := services.CreateDevice(ID)

	if err != nil {
		http.Error(w, "Error to add device", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(device)
}

func DeleteDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var deviceSelect struct {
		ID int `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&deviceSelect); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := services.DeleteDevice(deviceSelect.ID)

	if err != nil {
		http.Error(w, "Error to delete device", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("Device deleted successfully")
}
func GetLogsHandler(w http.ResponseWriter, r *http.Request) {
	var userSelect struct {
		ID int `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userSelect); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := services.GetUser(userSelect.ID)

	if err != nil {
		http.Error(w, "Error to get user", http.StatusInternalServerError)
		return
	}

	Logs, err := services.GetLogsByUser(user)
	if err != nil {
		http.Error(w, "Error to get logs", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(Logs)
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	var userSelect struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&userSelect); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := services.DeleteUser(userSelect.ID)

	if err != nil {
		http.Error(w, "Error to delete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetADSAllHandler(w http.ResponseWriter, r *http.Request) {

	media, err := services.GetAllMedia()
	if err != nil {
		http.Error(w, "Error to fetch users", http.StatusInternalServerError)
		return
	}

	host := r.Host
	var response []map[string]string
	for _, m := range media {
		adLink := fmt.Sprintf("http://%s%s", host, m.FilePath)

		response = append(response, map[string]string{
			"ID":   strconv.Itoa(m.ID),
			"name": m.MediaName,
			"link": adLink,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error to fetch users", http.StatusInternalServerError)
		return
	}
	return
}

func DeleteADSAllHandler(w http.ResponseWriter, r *http.Request) {
	var mediaSelect struct {
		ID int `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&mediaSelect); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	media, err := services.GetMediaByID(mediaSelect.ID)
	if err != nil {
		http.Error(w, "Error to fetch media", http.StatusInternalServerError)
		return
	}

	err = services.DeleteMedia(&media)
	if err != nil {
		http.Error(w, "Error to delete media", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//func UpdateDiveceHander(w http.ResponseWriter, r *http.Request) {
//	ID := r.Context().Value(UserIDKey).(int)
//	var device struct {
//		ID int `json:"id"`
//
//	}

//}
