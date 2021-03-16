package dataUser

type User struct {
	Id       int64  `json:"id" sql:"user_id"`
	Username string `json:"username" sql:"username"`
	Password string `json:"password" sql:"user_password"`
	Email    string `json:"email" sql:"email"`
}

type Users []*User
