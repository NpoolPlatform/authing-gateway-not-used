package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/authinggateway"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedAuthingGatewayServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterAuthingGatewayServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterAuthingGatewayHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
