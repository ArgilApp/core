package user

type UserSelfView struct {
	Id          uint64 `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type UserProfileView struct {
	Id          uint64 `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}
