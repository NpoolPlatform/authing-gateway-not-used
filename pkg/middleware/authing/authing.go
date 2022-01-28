package authing

import (
	"context"

	grpc2 "github.com/NpoolPlatform/authing-gateway/pkg/grpc"
	appusermgrpb "github.com/NpoolPlatform/message/npool/appusermgr"
	npool "github.com/NpoolPlatform/message/npool/authinggateway"
	logingwpb "github.com/NpoolPlatform/message/npool/logingateway"

	"golang.org/x/xerrors"
)

func AuthByApp(ctx context.Context, in *npool.AuthByAppRequest) (*npool.AuthByAppResponse, error) {
	resp, err := grpc2.GetApp(ctx, &appusermgrpb.GetAppRequest{
		ID: in.GetAppID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app: %v", err)
	}

	allowed := resp.Info != nil

	return &npool.AuthByAppResponse{
		Allowed: allowed,
	}, nil
}

func AuthByAppRoleUser(ctx context.Context, in *npool.AuthByAppRoleUserRequest) (*npool.AuthByAppRoleUserResponse, error) {
	resp, err := AuthByApp(ctx, &npool.AuthByAppRequest{
		AppID: in.GetAppID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail auth by app: %v", err)
	}
	if !resp.Allowed {
		return &npool.AuthByAppRoleUserResponse{
			Allowed: false,
		}, nil
	}

	resp1, err := grpc2.GetAppUser(ctx, &appusermgrpb.GetAppUserRequest{
		ID: in.GetUserID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app user: %v", err)
	}
	if resp1.Info == nil {
		return &npool.AuthByAppRoleUserResponse{
			Allowed: false,
		}, nil
	}

	_, err = grpc2.Logined(ctx, &logingwpb.LoginedRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetUserID(),
		Token:  in.GetToken(),
	})
	if err != nil {
		return nil, xerrors.Errorf("user not login: %v", err)
	}

	// TODO: check role access authorization to resource

	return nil, nil
}
