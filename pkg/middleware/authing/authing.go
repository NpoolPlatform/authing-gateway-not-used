package authing

import (
	"context"

	grpc2 "github.com/NpoolPlatform/authing-gateway/pkg/grpc"
	appusermgrpb "github.com/NpoolPlatform/message/npool/appusermgr"
	npool "github.com/NpoolPlatform/message/npool/authinggateway"

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
	return nil, nil
}
