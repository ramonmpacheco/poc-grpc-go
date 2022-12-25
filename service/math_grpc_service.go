package service

import (
	"context"

	"github.com/ramonmpacheco/poc-grpc-go/stub"
)

type mathGrpcService struct {
	// Gera compatibilidade com outras versões
	// Possibilida adicionar outro métodos depois
	stub.UnimplementedMathServiceServer
}

func NewMathGrpcService() *mathGrpcService {
	return &mathGrpcService{}
}

func (mgc *mathGrpcService) Sum(ctx context.Context, in *stub.SumRequest) (*stub.SumResponse, error) {
	result := in.Sum.A + in.Sum.B
	return &stub.SumResponse{
		Result: result,
	}, nil
}
