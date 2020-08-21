package user

import "fmt"

var users = []User{
	User{
		Id:          1,
		Username:    "test1",
		DisplayName: "Test User 1",
		Email:       "test1@argil.app",
	},
	User{
		Id:          2,
		Username:    "test2",
		DisplayName: "Test User 2",
		Email:       "test2@argil.app",
	},
	User{
		Id:          3,
		Username:    "test3",
		DisplayName: "Test User 3",
		Email:       "test3@argil.app",
	},
	User{
		Id:          4,
		Username:    "aaron",
		DisplayName: "Aaron Claydon",
		Email:       "aaron@argil.app",
	},
}

func GetUserByID(user *User, id uint64) (err error) {
	for i := 0; i < len(users); i++ {
		u := users[i]

		if u.Id == id {
			*user = u
			return nil
		}
	}

	return fmt.Errorf("Could not find a user with ID %d", id)
}

func GetUserByUsername(user *User, username string) (err error) {
	for i := 0; i < len(users); i++ {
		u := users[i]

		if u.Username == username {
			*user = u
			return nil
		}
	}

	return fmt.Errorf("Could not find a user with username %s", username)
}
