package main

import (
	"context"
	"log"
	"time"

	pb "productinfo/service/ordermgt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// 서버와의 연결을 구성한다

const address = "localhost:50052"


func main() {
	// 서버와의 연결을 구성한다.
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	orderMgtClient := pb.NewOrderManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 주문 정보 가져오기
	retrievedOrder, err := orderMgtClient.GetOrder(ctx, &wrapperspb.StringValue{Value: "106"})
	if err != nil {
		log.Fatalf("could not get order: %v", err)
	}
	log.Print("GetOrder Response -> : ", retrievedOrder)

}
