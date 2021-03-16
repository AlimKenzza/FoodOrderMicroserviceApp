package main

import (
	"context"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.com/AlimKenzza/authorization/pkg/dataUser"
	"gitlab.com/AlimKenzza/authorization/pkg/repository"
	"log"
	pbf "gitlab.com/AlimKenzza/authorization/userservice/userpb"
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
	dataUser.Conn, err = openDB(*dsn)
	if err != nil {
		log.Fatalf("Failed to connect to db: ", err)
	}
	userRepository = repository.NewUserRepository(dataUser.Conn)
	r := SetupRouter()
	r.Run(":4000")
}
