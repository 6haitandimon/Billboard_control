package app

import (
	"Billboard/internal/models"
	"Billboard/internal/repositories"
	"Billboard/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func AddAdminDeviceHandler(w http.ResponseWriter, r *http.Request) {
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

	host := os.Getenv("SERVER_API_HOST")
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

func UserDeviceListHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(UserIDKey).(int)

	devices, err := services.GetUserDevices(userID)

	if err != nil {
		http.Error(w, "Error to fetch user", http.StatusInternalServerError)
		json.NewEncoder(w).Encode(devices)
		return
	}

	json.NewEncoder(w).Encode(devices)

}

func GetFreeDevicesHandler(w http.ResponseWriter, r *http.Request) {
	devices, err := services.GetAllFreeDevices()

	if err != nil {
		http.Error(w, "Non free device", http.StatusBadGateway)
		return
	}

	json.NewEncoder(w).Encode(devices)

}

func AddUserDeviceHandler(w http.ResponseWriter, r *http.Request) {
	UserID := r.Context().Value(UserIDKey).(int)
	var deviceId struct {
		ID int `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&deviceId); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	device := models.Device{
		DeviceID:         deviceId.ID,
		UserID:           UserID,
		ConnectionStatus: false,
		LoadedAds:        "",
		GroupID:          1,
	}

	device, err := services.UpdateDevice(device)

	if err != nil {
		http.Error(w, "Error update device", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func GetDevicesGroupHandler(w http.ResponseWriter, r *http.Request) {
	UserID := r.Context().Value(UserIDKey).(int)

	deviceGroup, err := services.GetGroupDevice(UserID)

	if err != nil {
		json.NewEncoder(w).Encode(deviceGroup)
		http.Error(w, "Error to fetch group", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(deviceGroup)
}

func DeviceGroupCreateHandler(w http.ResponseWriter, r *http.Request) {
	UserID := r.Context().Value(UserIDKey).(int)
	var deviceGroup struct {
		Name string `json:"group_name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&deviceGroup); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}
	groupID, err := services.CreateGroup(deviceGroup.Name, UserID)

	if err != nil {
		json.NewEncoder(w).Encode(groupID)
		http.Error(w, "Error to create group", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(groupID)
}

func GetDeviceToGroupHandler(w http.ResponseWriter, r *http.Request) {
	UserID := r.Context().Value(UserIDKey).(int)

	var GroupId struct {
		ID int `json:"group_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&GroupId); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	device, err := services.GetUserDevicesByGroup(UserID, GroupId.ID)

	if err != nil {
		http.Error(w, "Error to fetch group", http.StatusInternalServerError)
		json.NewEncoder(w).Encode(device)
		return
	}
	json.NewEncoder(w).Encode(device)

}

func AddDevicesToGroupHandler(w http.ResponseWriter, r *http.Request) {
	UserID := r.Context().Value(UserIDKey).(int)
	var GroupId struct {
		GroupID  int `json:"group_id"`
		DeviceID int `json:"device_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&GroupId); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	device, err := services.AddToGroup(UserID, GroupId.DeviceID, GroupId.GroupID)
	if err != nil {
		json.NewEncoder(w).Encode(device)
		http.Error(w, "Error to add device to group", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(device)
}

func DeleteDevicesToGroupHandler(w http.ResponseWriter, r *http.Request) {
	//UserID := r.Context().Value(UserIDKey).(int)
	var GroupId struct {
		GroupID  int `json:"group_id"`
		DeviceID int `json:"device_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&GroupId); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := services.DeleteOnGroup(GroupId.DeviceID, GroupId.GroupID)
	if err != nil {
		//json.NewEncoder(w).Encode(device)
		http.Error(w, "Error to add device to group", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(device)
}

func GetScheduleHandler(w http.ResponseWriter, r *http.Request) {
	UserID := r.Context().Value(UserIDKey).(int)

	schedule, err := services.GetScheduleByUserID(UserID)

	if err != nil {
		http.Error(w, "Error to fetch schedule", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(schedule)
}

func ScheduleSettingHandler(w http.ResponseWriter, r *http.Request) {
	UserID := r.Context().Value(UserIDKey).(int)
	var scheduleID struct {
		ScheduleID int `json:"schedule_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&scheduleID); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	schedules, err := services.GetScheduleByUserID(UserID)
	var scheduleTrue models.Schedule
	if err != nil {
		http.Error(w, "Error to fetch schedule", http.StatusInternalServerError)
		return
	}

	for _, schedule := range schedules {
		if schedule.ID == scheduleID.ScheduleID {
			scheduleTrue = schedule
			break
		}
	}

	var scheduleFinale = models.ScheduleSender{
		ID:      scheduleTrue.ID,
		GroupID: scheduleTrue.GroupID,
		UserID:  scheduleTrue.UserID,
		Freq:    scheduleTrue.Freq,
	}

	adsId, err := services.DeserializeAdIDs(scheduleTrue.AdIDs)

	if err != nil {
		http.Error(w, "Error to deserialize ads", http.StatusInternalServerError)
		return
	}

	scheduleFinale.AdIDs = adsId

	json.NewEncoder(w).Encode(scheduleFinale)
}

//func ScheduleSettingSaveHandler(w http.ResponseWriter, r *http.Request) {
//	UserID := r.Context().Value(UserIDKey).(int)
//	var scheduleID struct {
//	}
//}

func GetDeviceADSHandler(w http.ResponseWriter, r *http.Request) {
	var DeviceID struct {
		DeviceID int `json:"device_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&DeviceID); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	group, err := services.GetDeviceByID(DeviceID.DeviceID)
	if err != nil {
		http.Error(w, "Error to fetch group", http.StatusInternalServerError)
		return
	}

	schedules, err := services.GetScheduleByUserID(group.UserID)
	if err != nil {
		http.Error(w, "Error to fetch schedule", http.StatusInternalServerError)
		return
	}

	var scheduleTrue models.Schedule
	for _, schedule := range schedules {
		if schedule.GroupID == group.GroupID {
			scheduleTrue = schedule
			break
		}
	}

	adsId, err := services.DeserializeAdIDs(scheduleTrue.AdIDs)

	if err != nil {
		http.Error(w, "Error to deserialize ads", http.StatusInternalServerError)
		return
	}
	fmt.Println(adsId)
	ads, err := services.GetMediaByID(adsId[0])
	if err != nil {
		http.Error(w, "Error to fetch ads", http.StatusInternalServerError)
		return
	}
	host := os.Getenv("SERVER_API_HOST")
	adLink := fmt.Sprintf("http://%s%s", host, ads.FilePath)

	ads.FilePath = adLink

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ads); err != nil {
		http.Error(w, "Error to fetch users", http.StatusInternalServerError)
		return
	}

}

func ScheduleSettingSaveHandler(w http.ResponseWriter, r *http.Request) {
	UserID := r.Context().Value(UserIDKey).(int)

}

func HelloPagesHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(w, "Listing...")
	//w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello Pages!"))
}

//func UpdateDiveceHander(w http.ResponseWriter, r *http.Request) {
//	ID := r.Context().Value(UserIDKey).(int)
//	var device struct {
//		ID int `json:"id"`
//
//	}

//}
