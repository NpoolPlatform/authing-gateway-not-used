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

func CreateGenesisAppUserAuth(ctx context.Context) ([]*npool.Auth, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.AuthingGatewayClient) (cruder.Any, error) {
		resp, err := cli.CreateGenesisAppUserAuth(ctx, &npool.CreateGenesisAppUserAuthRequest{})
		if err != nil {
			return nil, fmt.Errorf("fail notify email: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail notify email: %v", err)
	}

	return infos.([]*npool.Auth), nil
}

func GetAppAuths(ctx context.Context, appID string) ([]*npool.Auth, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.AuthingGatewayClient) (cruder.Any, error) {
		resp, err := cli.GetAuthsByApp(ctx, &npool.GetAuthsByAppRequest{
			AppID: appID,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get auth: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get auth: %v", err)
	}

	return infos.([]*npool.Auth), nil
}
