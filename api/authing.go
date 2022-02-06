package api

import (
	"context"

	appauthcrud "github.com/NpoolPlatform/authing-gateway/pkg/crud/appauth"
	authhistorycrud "github.com/NpoolPlatform/authing-gateway/pkg/crud/authhistory"
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
	resp, err := authhistorycrud.GetByAppUser(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get auth histories by app user: %v", err)
		return &npool.GetAuthHistoriesResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAuthHistoriesByApp(ctx context.Context, in *npool.GetAuthHistoriesByAppRequest) (*npool.GetAuthHistoriesByAppResponse, error) {
	resp, err := authhistorycrud.GetByApp(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get auth histories by app: %v", err)
		return &npool.GetAuthHistoriesByAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAuthHistoriesByOtherApp(ctx context.Context, in *npool.GetAuthHistoriesByOtherAppRequest) (*npool.GetAuthHistoriesByOtherAppResponse, error) {
	resp, err := authhistorycrud.GetByApp(ctx, &npool.GetAuthHistoriesByAppRequest{
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

func (s *Server) CreateAppAuthForOtherApp(ctx context.Context, in *npool.CreateAppAuthForOtherAppRequest) (*npool.CreateAppAuthForOtherAppResponse, error) {
	info := in.GetInfo()
	info.AppID = in.GetTargetAppID()
	resp, err := appauthcrud.CreateForOtherApp(ctx, &npool.CreateAppAuthForOtherAppRequest{
		Info: info,
	})
	if err != nil {
		logger.Sugar().Errorf("fail create app auth by other app: %v", err)
		return &npool.CreateAppAuthForOtherAppResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppAuthByAppResourceMethod(ctx context.Context, in *npool.GetAppAuthByAppResourceMethodRequest) (*npool.GetAppAuthByAppResourceMethodResponse, error) {
	resp, err := appauthcrud.GetAppAuthByAppResourceMethod(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail get app auth by app resource method: %v", err)
		return &npool.GetAppAuthByAppResourceMethodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) GetAppAuthByOtherAppResourceMethod(ctx context.Context, in *npool.GetAppAuthByOtherAppResourceMethodRequest) (*npool.GetAppAuthByOtherAppResourceMethodResponse, error) {
	resp, err := appauthcrud.GetAppAuthByAppResourceMethod(ctx, &npool.GetAppAuthByAppResourceMethodRequest{
		AppID:    in.GetTargetAppID(),
		Resource: in.GetResource(),
		Method:   in.GetMethod(),
	})
	if err != nil {
		logger.Sugar().Errorf("fail get app auth by other app resource method: %v", err)
		return &npool.GetAppAuthByOtherAppResourceMethodResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &npool.GetAppAuthByOtherAppResourceMethodResponse{
		Info: resp.Info,
	}, nil
}

func (s *Server) DeleteAppAuth(ctx context.Context, in *npool.DeleteAppAuthRequest) (*npool.DeleteAppAuthResponse, error) {
	resp, err := appauthcrud.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("fail delelp auth: %v", err)
		return &npool.DeleteAppAuthResponse{}, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) CreateAppRoleAuth(ctx context.Context, in *npool.CreateAppRoleAuthRequest) (*npool.CreateAppRoleAuthResponse, error) {
	return nil, nil
}

func (s *Server) CreateAppRoleAuthForOtherApp(ctx context.Context, in *npool.CreateAppRoleAuthForOtherAppRequest) (*npool.CreateAppRoleAuthForOtherAppResponse, error) {
	return nil, nil
}

func (s *Server) GetAppAuthByAppRoleResourceMethod(ctx context.Context, in *npool.GetAppAuthByAppRoleResourceMethodRequest) (*npool.GetAppAuthByAppRoleResourceMethodResponse, error) {
	return nil, nil
}

func (s *Server) GetAppAuthByOtherAppRoleResourceMethod(ctx context.Context, in *npool.GetAppAuthByOtherAppRoleResourceMethodRequest) (*npool.GetAppAuthByOtherAppRoleResourceMethodResponse, error) {
	return nil, nil
}

func (s *Server) DeleteAppRoleAuth(ctx context.Context, in *npool.DeleteAppRoleAuthRequest) (*npool.DeleteAppRoleAuthResponse, error) {
	return nil, nil
}

func (s *Server) CreateAppUserAuth(ctx context.Context, in *npool.CreateAppUserAuthRequest) (*npool.CreateAppUserAuthResponse, error) {
	return nil, nil
}

func (s *Server) CreateAppUserAuthForOtherApp(ctx context.Context, in *npool.CreateAppUserAuthForOtherAppRequest) (*npool.CreateAppUserAuthForOtherAppResponse, error) {
	return nil, nil
}

func (s *Server) GetAppAuthByAppUserResourceMethod(ctx context.Context, in *npool.GetAppAuthByAppUserResourceMethodRequest) (*npool.GetAppAuthByAppUserResourceMethodResponse, error) {
	return nil, nil
}

func (s *Server) GetAppAuthByOtherAppUserResourceMethod(ctx context.Context, in *npool.GetAppAuthByOtherAppUserResourceMethodRequest) (*npool.GetAppAuthByOtherAppUserResourceMethodResponse, error) {
	return nil, nil
}

func (s *Server) DeleteAppUserAuth(ctx context.Context, in *npool.DeleteAppUserAuthRequest) (*npool.DeleteAppUserAuthResponse, error) {
	return nil, nil
}

func (s *Server) GetAuthsByApp(ctx context.Context, in *npool.GetAuthsByAppRequest) (*npool.GetAuthsByAppResponse, error) {
	return nil, nil
}

func (s *Server) GetAuthsByOtherApp(ctx context.Context, in *npool.GetAuthsByOtherAppRequest) (*npool.GetAuthsByOtherAppResponse, error) {
	return nil, nil
}

func (s *Server) GetAuthsByAppRole(ctx context.Context, in *npool.GetAuthsByAppRoleRequest) (*npool.GetAuthsByAppRoleResponse, error) {
	return nil, nil
}

func (s *Server) GetAuthsByOtherAppRole(ctx context.Context, in *npool.GetAuthsByOtherAppRoleRequest) (*npool.GetAuthsByOtherAppRoleResponse, error) {
	return nil, nil
}

func (s *Server) GetAuthsByAppUser(ctx context.Context, in *npool.GetAuthsByAppUserRequest) (*npool.GetAuthsByAppUserResponse, error) {
	return nil, nil
}

func (s *Server) GetAuthsByOtherAppUser(ctx context.Context, in *npool.GetAuthsByOtherAppUserRequest) (*npool.GetAuthsByOtherAppUserResponse, error) {
	return nil, nil
}
