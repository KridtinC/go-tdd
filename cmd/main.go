package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kridtinc/go-tdd/internal/app/endpoint"
	"github.com/kridtinc/go-tdd/internal/app/usecase"
	"github.com/kridtinc/go-tdd/proto/userpb"
	"google.golang.org/grpc"
)

// PORT grpc server port
const PORT = "8080"

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", PORT))
	if err != nil {
		log.Fatal(err)
	}

	userUseCase := usecase.NewUserUseCase()
	userEndpoint := endpoint.NewUserEndpoint(userUseCase)
	grpcServ := grpc.NewServer()

	userpb.RegisterUserServer(grpcServ, userEndpoint)

	log.Printf("listen at grpc server :%s", PORT)

	if err := grpcServ.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
