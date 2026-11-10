package main

import (
	"context"
	"log"
	"time"
	pb "productinfo/client/ecommerce" // 프로토버프 컴파일러로 생성한 코드가 포함된 패키지를 임포트함
	"google.golang.org/grpc" 
)

const address := "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // 제공된 주소(localhost 50051)로 서바와의 커넥션을 설정함, 
														// 여기서는 클라이언트와 서버 사이에 보안되지 않은 커넥션을 만듦
	if err != nil{
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close() // 모든 작업이 끝나면 커넥션을 종료한다.
	c := pb.NewProductInfoClient(conn) // 커넥션을 전달해 스텁을 생성함 이 스텁 인스턴스는 서버를 호출하는 모든 원격 메서드가 포함돼 있음

	name := "Apple iPhone 11"
	description := `Meet Apple iPhone 11 All-new dual-camera system with Ultra Wide and Night mode.`
	price := float32(1000.0)

	ctx, cancel := context.WithTimeout(context.Background(), time.second) // 원격 호출과 함께 전달할 Context를 생성함, 
																		// 이 Context 객체에는 최종 사용자 인증 토큰에 대한 ID와 요청 데드라인 같은 메타데이터가 포함되며,
																		// 서비스 요청 동안 유지됨

	defer cancel()
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, description: description, Price: price}) // 제품 정보와 함께 addProduct 메서드를 호출한다.
																								// 호출이 완료되면 제품 정보를 반환하고, 그렇지 않으면 에러가 반환된다.
	if != nil{
		log.Fatalf("copuld not get product: %v", err)
	}

	log.Printf("product: ", product.String())
}