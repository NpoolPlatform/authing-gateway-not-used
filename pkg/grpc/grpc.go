package grpc

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	appusermgrconst "github.com/NpoolPlatform/appuser-manager/pkg/message/const" //nolint
	appusermgrpb "github.com/NpoolPlatform/message/npool/appusermgr"

	"golang.org/x/xerrors"
)

const (
	grpcTimeout = 5 * time.Second
)

//---------------------------------------------------------------------------------------------------------------------------

func GetApp(ctx context.Context, in *appusermgrpb.GetAppRequest) (*appusermgrpb.GetAppResponse, error) {
	conn, err := grpc2.GetGRPCConn(appusermgrconst.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, xerrors.Errorf("fail get app user connection: %v", err)
	}
	defer conn.Close()

	cli := appusermgrpb.NewAppUserManagerClient(conn)

	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	return cli.GetApp(ctx, in)
}
