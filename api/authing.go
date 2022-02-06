package api

import (
	"context"

	crud "github.com/NpoolPlatform/authing-gateway/pkg/crud/authhistory"
	mw "github.com/NpoolPlatform/authing-gateway/pkg/middleware/authing"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/authinggateway"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AuthByApp(ctx context.Context, in *npool.AuthByAppRequest) (*npool.AuthByAppResponse, error) {
	resp, err := mw.AuthByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail auth by app: %v", err)
		return &npool.AuthByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) AuthByAppRoleUser(ctx context.Context, in *npool.AuthByAppRoleUserRequest) (*npool.AuthByAppRoleUserResponse, error) {
	resp, err := mw.AuthByAppRoleUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail auth by app role user: %v", err)
		return &npool.AuthByAppRoleUserResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAuthHistories(ctx context.Context, in *npool.GetAuthHistoriesRequest) (*npool.GetAuthHistoriesResponse, error) {
	resp, err := crud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get auth histories by app user: %v", err)
		return &npool.GetAuthHistoriesResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAuthHistoriesByApp(ctx context.Context, in *npool.GetAuthHistoriesByAppRequest) (*npool.GetAuthHistoriesByAppResponse, error) {
	resp, err := crud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get auth histories by app: %v", err)
		return &npool.GetAuthHistoriesByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAuthHistoriesByOtherApp(ctx context.Context, in *npool.GetAuthHistoriesByOtherAppRequest) (*npool.GetAuthHistoriesByOtherAppResponse, error) {
	resp, err := crud.GetByApp(ctx, &npool.GetAuthHistoriesByAppRequest{
		AppID: in.GetTargetAppID(),
	})
	if err != nil {
		logger.Sugar().Errorf("fail get auth histories by other app: %v", err)
		return &npool.GetAuthHistoriesByOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAuthHistoriesByOtherAppResponse{
		Infos: resp.Infos,
	}, nil
}
