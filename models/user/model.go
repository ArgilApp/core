package user

type User struct {
	Id           uint64 `json:"id"`
	Username     string `json:"username"`
	DisplayName  string `json:"display_name"`
	Email        string `json:"email"`
	PasswordHash string
}

func (b *User) TableName() string {
	return "users"
}
