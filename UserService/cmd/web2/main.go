package main

import (
	"context"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.com/AlimKenzza/authorization/pkg/dataUser"
	"gitlab.com/AlimKenzza/authorization/pkg/repository"
	"log"
	pbf "gitlab.com/AlimKenzza/authorization/userpb"
)

func openDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Println("Connection for database is established")
		return nil, err
	}
	return pool, nil
}

type IRepository interface{
	GetUser(request *pbf.UserGetRequest) *pbf.UserGetResponse
}
type Repository struct{
	dataUser.User
}
type service struct {
	repo IRepository
}
func (s *service) GetUser(tx context.Context, req *pbf.UserGetRequest)(*pbf.UserGetResponse, error){
	user := s.repo.GetUser(req)
	return &pbf.UserGetResponse{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	},nil
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

