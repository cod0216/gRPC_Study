package main

import (
	"log"
	"net"

	pb "productinfo/service/ecommerce" // 프로토버프 컴파일러로 생성된 코드가 포함된 패키지를 임포트한

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50051"

func main() {
	lis, err := net.Listen("tcp", port) // gRPC 서버가 바인딩하고자 하는 TCP 리스너는 지정 포트(50051)로 생성된다.
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() // RPC Go API를 호출해 새 gRPC 서버 인스턴스를 생성
	pb.RegisterProductInfoServer(s, &server{}) // 이전에 구현된 서비스로 생성된 API를 호출해 새로 작성된 gRPC 서버에 등록한다.
	reflection.Register(s)

	log.Printf("starting gRPC listener on port %s", port)
	if err := s.Serve(lis); err != nil {


		log.Fatalf("failed to serve: %v", err)
	} // 50051에서 수신되는 메시지를 리스닝 하기 시작
}
