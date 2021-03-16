package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.com/AlimKenzza/authorization/interfaces"
	"gitlab.com/AlimKenzza/authorization/pkg/dataOrders"
	"log"
)

type OrderRepository struct{
	pool pgxpool.Pool
}

func (o OrderRepository) GetAllOrders() []*dataOrders.Order {
	stmt:= "SELECT * FROM Orders"
	rows, err := o.pool.Query(context.Background(),stmt)
	if err != nil {
		log.Fatal("Failed to SELECT: %v", err)
		return nil
	}
	defer rows.Close()
	orders := []*dataOrders.Order{}
	for rows.Next() {
		o := &dataOrders.Order{}
		err = rows.Scan(&o.ID, &o.UserID, &o.FoodID, &o.Count)
		if err != nil {
			log.Fatalf("Failed to scan: %v", err)
			return nil
		}
		orders = append(orders, o)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return orders
}

func (r OrderRepository) GetOrderById(id int) *dataOrders.Order{
	stmt := "SELECT * FROM orders WHERE order_id = $1"
	o := &dataOrders.Order{}
	err := r.pool.QueryRow(context.Background(), stmt, id).Scan(&o.ID, &o.UserID, &o.FoodID, &o.Count)
	if err != nil {
		log.Println("Didn't find order with id ", id)
		return nil
	}
	return o
}

func (o OrderRepository) DeleteOrder(order dataOrders.Order) bool {
	_, err := o.pool.Exec(context.Background(),
		"DELETE FROM orders WHERE order_id = $1", order.ID)
	if err != nil {
		return false
	}
	return true
}

func (o OrderRepository) CreateOrder(order dataOrders.Order) bool {
	_, err := o.pool.Exec(context.Background(),
		"INSERT INTO ORDERS VALUES($1,$2,$3,$4)",
		order.ID, order.UserID, order.FoodID, order.Count)
	if err != nil {
		return false
	}
	return true
}

func NewOrderRepository(conn *pgxpool.Pool) interfaces.IOrderRepository{
	return &OrderRepository{*conn}
}