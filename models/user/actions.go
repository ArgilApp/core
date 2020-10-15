package user

import (
	"fmt"

	"github.com/argilapp/core/db"
)

func GetUserByID(id uint64) (*User, error) {
	var user User

	if result := db.Session.First(&user, id); result.Error != nil {
		return nil, fmt.Errorf("Could not find a user with ID %d", id)
	}

	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	var user User

	if result := db.Session.Where("username = ?", username).First(&user); result.Error != nil {
		return nil, fmt.Errorf("Could not find a user with username %s", username)
	}

	return &user, nil
}
