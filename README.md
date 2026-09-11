# gRPC를 공부하는 엄숙한 공간 입니다.

이곳은 gRPC를 공부 하기 위한 프라이빗하고 
**아주아주아주아주 엄숙한 공간**입니다.

# 이유

채소 가게에 일하게 됐는데 그녀석이 gRPC를 쓰고 있어서 공부해야 합니다.

하다가 어지러우면 그녀석을 흔들 수 도 있습니다.(발견 즉시 살려주세요)


# 목표

```
Protocol Buffers / IDL
      ↓
protoc / 코드 생성
      ↓
Unary RPC — 서버와 클라이언트
      ↓
스트리밍 (Server / Client / Bidirectional)
      ↓
메타데이터 / 에러 코드 / 인터셉터
      ↓
Deadline / 취소 / 재시도
      ↓
테스트 (bufconn) / grpcurl
      ↓
같은 계약을 다른 언어로 (Go ↔ Java)
```

gRPC의 사용법을 아는 것뿐만 아니라,

> **왜 gRPC는 이렇게 동작하는가?**

를 이해하는 것을 목표로 합니다.
