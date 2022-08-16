package models

type User struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Firstname  string `json:"firstName"`
	Lastname   string `json:"lastName"`
	Email      string `json:"email"`
	Pass       string `json:"password"`
	Phone      string `json:"phone"`
	Userstatus int64  `json:"userStatus"`
}
