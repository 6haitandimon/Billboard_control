package services

import (
	"Billboard/internal/models"
	"Billboard/internal/repositories"
	"Billboard/pkg/auth"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Authenticate(username, password string) (string, error, int) {
	user, err := repositories.GetUserByName(username)
	if err != nil {
		return "", errors.New("user not found"), -1
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid password"), -1
	}

	if user.RoleID == 1 {
		token, err := auth.GenerateToken(user.ID, "user")
		if err != nil {
			return "", errors.New("failed to generate token"), -1
		}
		return token, nil, user.RoleID
	} else if user.RoleID == 2 {
		token, err := auth.GenerateToken(user.ID, "admin")
		if err != nil {
			return "", errors.New("failed to generate token"), -1
		}
		return token, nil, user.RoleID
	}

	return "", errors.New("invalid password"), -1
}

func Registration(username string, password string, role int) error {
	passHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	newUser := models.User{
		FullName:         "",
		Phone:            "",
		RegistrationDate: time.Now().Format("2006-01-02"),
		DeviceSerial:     "",
		PasswordHash:     string(passHash),
		RoleID:           1,
		UserName:         username,
	}

	if username == "admin" {
		newUser.RoleID = 2
	}

	err := repositories.AddUser(&newUser)

	if err != nil {
		return err
	}
	return nil
}
