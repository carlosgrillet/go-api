package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/carlosgrillet/go-api/db"
	"github.com/carlosgrillet/go-api/utils"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

func (u *User) Save() error {
  u.Password,_ = utils.Encrypt(u.Password)
	jsonData, err := json.Marshal(u)
	if err != nil {
		return errors.New("Error encoding the user to json")
	}

	db.Put("/users/", u.ID, string(jsonData))
	return nil
}

func (u *User) Delete() error {
	err := db.Delete("/users/", u.ID)
	if err != nil {
		return errors.New("Failed to delete event")
	}
	return nil
}

func (u *User) ValidateCredentials() (string, error) {
  user, err := GetUserByEmail(u.Email)
  if err != nil {
    return "", errors.New("User not registred")
  }

  if !utils.CompareWithHash(u.Password, user.Password) {
    return "", errors.New("Bad credentials")
  }
  return user.ID, nil
}

func GetUserById(id string) (User, error) {
	userInDb, err := db.Get("/users/", id, false)
	if err != nil {
		return User{}, errors.New("User not found")
	}

	var user User
	for _, value := range userInDb {
		json.Unmarshal([]byte(value), &user)
	}
	return user, nil
}

func GetUserByEmail(email string) (User, error) {
  userInDb, err := db.Get("/users/", "", true)
	if err != nil {
		return User{}, errors.New("User not found")
	}

	var user User
	var searchUser User
	for _, value := range userInDb {
		json.Unmarshal([]byte(value), &user)
    if user.Email == email{
      json.Unmarshal([]byte(value), &searchUser)
      break
    }
	}
  if searchUser.Email == "" {
		return User{}, errors.New("User not found")
  }
	return searchUser, nil
}
