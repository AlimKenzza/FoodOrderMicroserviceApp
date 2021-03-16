package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.com/AlimKenzza/authorization/foodService"
	"gitlab.com/AlimKenzza/authorization/pkg/dataFood"
	"log"
)

type FoodRepository struct {
	pool pgxpool.Pool
}

func NewFoodRepository(conn *pgxpool.Pool) foodService.IFoodsRepository {
	return &FoodRepository{*conn}
}

func (r *FoodRepository) GetFoodById(id int) *dataFood.Food {
	stmt := "SELECT * FROM food WHERE food_id = $1"
	food := &dataFood.Food{}
	err := r.pool.QueryRow(context.Background(), stmt, id).Scan(&food.ID, &food.FoodName, &food.Description, &food.Quantity, &food.Price)
	if err != nil {
		log.Println("Didn't find food with id ", id)
		return nil
	}
	return food
}

func (r FoodRepository) GetAllFoods() []*dataFood.Food {
	stmt := "SELECT * FROM food"
	rows, err := r.pool.Query(context.Background(), stmt)
	if err != nil {
		log.Fatal("Failed to SELECT: %v", err)
		return nil
	}
	defer rows.Close()
	foods := []*dataFood.Food{}
	for rows.Next() {
		o := &dataFood.Food{}
		err = rows.Scan(&o.ID, &o.FoodName, &o.Description, &o.Quantity, &o.Price)
		if err != nil {
			log.Fatalf("Failed to scan: %v", err)
			return nil
		}
		foods = append(foods, o)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return foods
}

func (r *FoodRepository) DeleteFood(food dataFood.Food) bool {
	_, err := r.pool.Exec(context.Background(),
		"DELETE FROM food WHERE food_id = $1", food.ID)
	if err != nil {
		return false
	}
	return true
}

func (r *FoodRepository) CreateFood(food dataFood.Food) bool {
	sql := "INSERT INTO food(food_name, description, quantity, price) " +
		"VALUES($1, $2, $3, $4) RETURNING food_id"
	row := r.pool.QueryRow(context.Background(),
		sql, food.FoodName, food.Description, food.Quantity, food.Price)
	var id int
	err := row.Scan(&id)
	if err != nil {
		log.Printf("Unable to INSERT: %v\n", err)
		return false
	}
	return true
}
