package main

import (
	"context"
	"fmt"
	restarauntpb "gitlab.com/AlimKenzza/authorization"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"os"
	"strconv"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	args := os.Args
	conn, err := grpc.Dial("0.0.0.0:4001", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	orderId, err := strconv.Atoi(args[1])
	client := restarauntpb.NewFoodOrderServiceClient(conn)
	request := &restarauntpb.FoodsByOrderRequest{OrderID: int32(orderId)}
	response, err := client.GetOFoodsByOrder(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response.Context())
}
