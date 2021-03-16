package data

const (
	CheckUserExists = `SELECT user_id from users WHERE username = $1`
	LoginQuery      = `SELECT * from users WHERE username = $1`
	CreateUserQuery = `INSERT INTO users(user_id,username,user_password,email) VALUES (DEFAULT, $1 , $2, $3);`
)
