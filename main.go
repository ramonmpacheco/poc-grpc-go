package main

import (
	"log"
	"net"

	controllergrpc "github.com/ramonmpacheco/poc-grpc-go/service"
	"github.com/ramonmpacheco/poc-grpc-go/stub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	grpcServer := grpc.NewServer()
	stub.RegisterMathServiceServer(grpcServer, controllergrpc.NewMathGrpcService())
	// para conseguir ler e processar sua propria informação
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("cannot start server. err=%v", err)
	}

	log.Default().Println("GRPC Server will be running at port: 50051")

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start server. err=%v", err)
	}
}
