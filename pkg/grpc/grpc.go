package grpc

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	appusermgrconst "github.com/NpoolPlatform/appuser-manager/pkg/message/const" //nolint
	appusermgrpb "github.com/NpoolPlatform/message/npool/appusermgr"

	logingwconst "github.com/NpoolPlatform/login-gateway/pkg/message/const"
	logingwpb "github.com/NpoolPlatform/message/npool/logingateway"

	"golang.org/x/xerrors"
)

const (
	grpcTimeout = 5 * time.Second
)

//---------------------------------------------------------------------------------------------------------------------------

func GetAppInfo(ctx context.Context, in *appusermgrpb.GetAppInfoRequest) (*appusermgrpb.GetAppInfoResponse, error) {
	conn, err := grpc2.GetGRPCConn(appusermgrconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get app user connection: %v", err)
	}
	defer conn.Close()

	cli := appusermgrpb.NewAppUserManagerClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	return cli.GetAppInfo(ctx, in)
}

func GetAppUserInfo(ctx context.Context, in *appusermgrpb.GetAppUserInfoRequest) (*appusermgrpb.GetAppUserInfoResponse, error) {
	conn, err := grpc2.GetGRPCConn(appusermgrconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get app user connection: %v", err)
	}
	defer conn.Close()

	cli := appusermgrpb.NewAppUserManagerClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	return cli.GetAppUserInfo(ctx, in)
}

func GetGenesisAppRoleUsersByOtherApp(ctx context.Context, in *appusermgrpb.GetGenesisAppRoleUsersByOtherAppRequest) (*appusermgrpb.GetGenesisAppRoleUsersByOtherAppResponse, error) {
	conn, err := grpc2.GetGRPCConn(appusermgrconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get app user connection: %v", err)
	}
	defer conn.Close()

	cli := appusermgrpb.NewAppUserManagerClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	return cli.GetGenesisAppRoleUsersByOtherApp(ctx, in)
}

//---------------------------------------------------------------------------------------------------------------------------

func Logined(ctx context.Context, in *logingwpb.LoginedRequest) (*logingwpb.LoginedResponse, error) {
	conn, err := grpc2.GetGRPCConn(logingwconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get login gateway connection: %v", err)
	}
	defer conn.Close()

	cli := logingwpb.NewLoginGatewayClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	return cli.Logined(ctx, in)
}
