package client

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/authing-gateway/pkg/message/const"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/authinggateway"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.AuthingGatewayClient) (cruder.Any, error)) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get third gateway connection: %v", err)
	}
	defer conn.Close()

	cli := npool.NewAuthingGatewayClient(conn)

	return fn(_ctx, cli)
}

func CreateGenesisAppUserAuth(ctx context.Context, in *npool.CreateGenesisAppUserAuthRequest) (*npool.CreateGenesisAppUserAuthResponse, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.AuthingGatewayClient) (cruder.Any, error) {
		resp, err := cli.CreateGenesisAppUserAuth(ctx, in)
		if err != nil {
			return nil, fmt.Errorf("fail notify email: %v", err)
		}
		return resp, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail notify email: %v", err)
	}

	return info.(*npool.CreateGenesisAppUserAuthResponse), nil
}

func GetAuthsByOtherApp(ctx context.Context, appID string) ([]*npool.Auth, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.AuthingGatewayClient) (cruder.Any, error) {
		resp, err := cli.GetAuthsByOtherApp(ctx, &npool.GetAuthsByOtherAppRequest{
			TargetAppID: appID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get auth: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get auth: %v", err)
	}

	return info.([]*npool.Auth), nil
}