package interfaces

import "gitlab.com/AlimKenzza/authorization/pkg/dataUser"

type IUsersRepository interface {
	//CreateOrder(user dataUser.User) bool
	GetAllUsers() []*dataUser.User
	GetUserById(id int) *dataUser.User
	DeleteUser(order dataUser.User) bool
	UpdateUser(order dataUser.User) bool
}
