package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/AlimKenzza/authorization/interfaces"
	"strconv"
)

const (
	queryGetUser                = "SELECT * FROM users WHERE user_id=$1"
	queryUpdateUser             = "UPDATE users SET username=$1, email=$2 WHERE user_id=$3"
	queryDeleteUser             = "DELETE FROM users WHERE user_id=$1"
	queryFindByEmailAndPassword = "SELECT user_id, username, email FROM users WHERE email=$1 AND user_password=$2"
)

var jsonContentType = "application/json; charset=utf-8"
var userRepository interfaces.IUsersRepository

func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	order := userRepository.GetUserById(id)
	c.JSON(200, order)
}

func GetAllUsers(c *gin.Context) {
	users := userRepository.GetAllUsers()
	c.JSON(200, users)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.Data(400, jsonContentType, []byte("Incorrect id format"))
		return
	}
	user := userRepository.GetUserById(id)
	if user == nil {
		c.Data(400, jsonContentType, []byte("No such user with id"))
		return
	}
	if userRepository.DeleteUser(*user) {
		c.Data(200, jsonContentType, []byte("User deleted successfully"))
		return
	}
	c.Data(500, jsonContentType, []byte("Failed to delete user"))
}

//func UpdateOrder(c *gin.Context)  {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil || id < 1 {
//		c.Data(400, jsonContentType, []byte("Incorrect id"))
//		return
//	}
//	model := userRepository.GetUserById(id)
//	user := &dataUser.User{}
//	err = c.BindJSON(user)
//	if err != nil {
//		c.Data(400, jsonContentType, []byte("Fill all fields"))
//		return
//	}
//	user.Id = int64(id)
//	updateValues(model, user)
//	if userRepository.UpdateUser(*user) {
//		c.Data(200, jsonContentType, []byte("Updated user"))
//		return
//	}
//	c.Data(500, jsonContentType, []byte("Failed to update user"))
//}
