package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	restarauntpb "gitlab.com/AlimKenzza/authorization"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0"+os.Getenv("HOST"))
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	restarauntpb.RegisterFoodOrderServiceServer(s, &Server{})
	log.Println("Server is running on port ", os.Getenv("HOST"))
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}

type Server struct{
	restarauntpb.UnimplementedFoodOrderServiceServer
}
func (s *Server) GetFoods (req *restarauntpb.FoodsByOrderRequest, stream restarauntpb.FoodOrderService_GetOFoodsByOrderServer) error {
	orderId := req.GetOrderID()
	foods := requestDispatcher(orderId)
	for i:=0;i<len(foods);i++{
		response := &restarauntpb.FoodsByOrderResponse{
			FoodID:      int32(foods[i].ID),
			FoodName:    foods[i].FoodName,
			Description: foods[i].Description,
			Quantity:    int32(foods[i].Quantity),
			Price:       int32(foods[i].Price),
		}
		err := stream.Send(response)
		if err != nil {
			log.Printf("Error while sending %v", err)
		}
	}
	return nil

}
func requestDispatcher(orderID int32) ([]*restarauntpb.Food){
	var pool *pgxpool.Pool
	var err error
	pool, err = OpenDB(os.Getenv("CONN"))
	if err != nil {
		log.Fatalf("Failed to connect to db: ", err)
	}
	var foods []*restarauntpb.Food
	foods = findFoodsByOrderId(pool,orderID)
	return foods
}
func findFoodsByOrderId(pool *pgxpool.Pool,orderID int32) ([]*restarauntpb.Food){
	rows, err := pool.Query(context.Background(),"SELECT * FROM food JOIN orders ON food.food_id = orders.food_id WHERE order_id=$1",orderID)
	if err != nil {
		log.Printf("BAD GET REQUEST: %v", err)
		log.Printf("Didn't found order")
		return nil
	}
	defer rows.Close()
	var foods []*restarauntpb.Food
	for rows.Next(){
		food := &restarauntpb.Food{}
		err = rows.Scan(&food.ID,&food.FoodName,&food.Description,&food.Price,&food.Quantity)
		if err != nil {
			return nil
		}
		foods = append(foods,food)
	}
	return foods
}
func OpenDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Println("Connection for database couldn't be established")
		return nil, err
	}
	return pool, nil
}