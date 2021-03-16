package dataFood

type Food struct {
	ID          int    `json:"id" sql:"food_id"`
	FoodName    string `json:"foodName" sql:"food_name"`
	Description string `json:"description" sql:"description"`
	Quantity    int    `json:"quantity" sql:"quantity"`
	Price       int    `json:"price" sql:"price"`
}
