package controllergrpc

import (
	"context"

	"github.com/ramonmpacheco/poc-grpc-go/stub"
)

type mathGrpcController struct {
	stub.UnimplementedMathServiceServer
}

func NewMathGrpcController() *mathGrpcController {
	return &mathGrpcController{}
}

func (mgc *mathGrpcController) Sum(ctx context.Context, in *stub.SumRequest) (*stub.SumResponse, error) {
	result := in.Sum.GetA() + in.Sum.GetB()
	return &stub.SumResponse{
		Result: result,
	}, nil
}
