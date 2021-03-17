package main

import (
	"gitlab.com/AlimKenzza/authorization/pkg/data"
	"gitlab.com/AlimKenzza/authorization/router"
)

func init() {
	data.Connect()
}

func main() {
	r := router.SetupRouter()
	// Listen and Serve in 0.0.0.0:8081
	r.Run(":8081")
	//dsn := flag.String("dsn", "postgresql://localhost/restaurant?user=postgres&password=alimzhan125", "PostGreSQL")
	//flag.Parse()
	//var err error
	//data.Conn, err = OpenDB(*dsn)
	//if err != nil {
	//	log.Fatalf("Failed to connect to db: ", err)
	//}
	//handlers.UserRepository = repositories.NewUserRepository(data.Conn)
	//r := router.SetupRouter()
	//r.Run(":8081")
}

//func OpenDB(dsn string) (*pgxpool.Pool, error) {
//	pool, err := pgxpool.Connect(context.Background(), dsn)
//	if err != nil {
//		log.Println("Connection for database couldn't be established")
//		return nil, err
//	}
//	return pool, nil
//}
