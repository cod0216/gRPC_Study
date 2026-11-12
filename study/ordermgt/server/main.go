package main

import (
        "context"
        "log"
        "net"

        pb "productinfo/service/ordermgt"

        "google.golang.org/grpc"
        "google.golang.org/grpc/codes"
        "google.golang.org/grpc/status"
        "google.golang.org/protobuf/types/known/wrapperspb"
)

const port = ":50052"

var orderMap = map[string]*pb.Order{
        "102": {Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 1800.00},
}

type server struct {
        pb.UnimplementedOrderManagementServer
}

func (s *server) GetOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error) {
        ord, exists := orderMap[orderId.Value]
        if !exists {
                return nil, status.Errorf(codes.NotFound, "Order does not exist: %s", orderId.Value)
        }
        return ord, nil
}

func main() {
        lis, err := net.Listen("tcp", port)
        if err != nil {
                log.Fatalf("failed to listen: %v", err)
        }
        s := grpc.NewServer()
        pb.RegisterOrderManagementServer(s, &server{})
        log.Printf("Starting gRPC listener on port %s", port)
        if err := s.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
        }
}