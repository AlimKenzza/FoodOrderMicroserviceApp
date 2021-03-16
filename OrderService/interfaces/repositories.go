package interfaces

import "gitlab.com/AlimKenzza/authorization/pkg/dataOrders"

type IOrderRepository interface{
	GetAllOrders() []*dataOrders.Order
	GetOrderById(int int) *dataOrders.Order
	DeleteOrder(order dataOrders.Order) bool
	CreateOrder(order dataOrders.Order) bool
}
