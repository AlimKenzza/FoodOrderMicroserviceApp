package dataOrders

type Order struct {
	ID          int    `json:"id" sql:"order_id"`
	UserID    int `json:"user_id" sql:"user_id"`
	FoodID int `json:"food_id" sql:"food_id"`
	Count    int    `json:"count" sql:"count"`
}