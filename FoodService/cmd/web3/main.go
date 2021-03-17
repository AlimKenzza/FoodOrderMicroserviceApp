package main

import (
	"context"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.com/AlimKenzza/authorization/pkg/dataFood"
	"gitlab.com/AlimKenzza/authorization/pkg/repository"
	"log"
)

func openDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Println("Connection for database is established")
		return nil, err
	}
	return pool, nil
}

func main() {
	dsn := flag.String("dsn", "postgresql://localhost/restaurant?user=postgres&password=alimzhan125", "PostGreSQL")
	flag.Parse()
	var err error
	dataFood.Conn, err = openDB(*dsn)
	if err != nil {
		log.Fatalf("Failed to connect to db: ", err)
	}
	foodRepository = repository.NewFoodRepository(dataFood.Conn)
	r := SetupRouter()
	r.Run(":4040")
}
