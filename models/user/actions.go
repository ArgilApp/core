package user

import (
	"fmt"
	"log"

	"github.com/alexedwards/argon2id"
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

func Authenticate(username string, password string) bool {
	user, err := GetUserByUsername(username)

	if err != nil {
		return false
	}

	passwordMatch := user.VerifyPassword(password)

	return passwordMatch
}

func (u *User) SetPassword(password string) error {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)

	if err != nil {
		return err
	}

	u.PasswordHash = hash

	return nil
}

func (u *User) VerifyPassword(password string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, u.PasswordHash)

	if err != nil {
		log.Println("Failed verifying password", err)
		return false
	}

	return match
}
