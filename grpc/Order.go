package restarauntpb

type Food struct {
	ID          int
	FoodName    string
	Description string
	Quantity    int
	Price       int
}
type Order struct {
	ID          int
	UserID    int
	FoodID int
	Count    int
}