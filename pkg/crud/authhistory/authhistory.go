package authhistory

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/authing-gateway/pkg/db"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/authhistory"
	npool "github.com/NpoolPlatform/message/npool/authinggateway"

	"github.com/google/uuid"
)

const (
	grpcTimeout = 5 * time.Second
)

func Create(ctx context.Context, info *npool.AuthHistory) error {
	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return fmt.Errorf("fail get db client: %v", err)
	}

	appID, err := uuid.Parse(info.GetAppID())
	if err != nil {
		return fmt.Errorf("invalid app id: %v", err)
	}

	userID, err := uuid.Parse(info.GetUserID())
	if err != nil {
		userID = uuid.UUID{}
	}

	_, err = cli.
		AuthHistory.
		Create().
		SetAppID(appID).
		SetUserID(userID).
		SetResource(info.GetResource()).
		SetMethod(info.GetMethod()).
		SetAllowed(info.GetAllowed()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("fail create auth history: %v", err)
	}

	return nil
}

func GetByAppUser(ctx context.Context, in *npool.GetAuthHistoriesRequest) (*npool.GetAuthHistoriesResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	userID, err := uuid.Parse(in.GetUserID())
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %v", err)
	}

	infos, err := cli.
		AuthHistory.
		Query().
		Where(
			authhistory.And(
				authhistory.AppID(appID),
				authhistory.UserID(userID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query auth history: %v", err)
	}

	ahs := []*npool.AuthHistory{}
	for _, info := range infos {
		ahs = append(ahs, &npool.AuthHistory{
			ID:       info.ID.String(),
			AppID:    info.AppID.String(),
			UserID:   info.UserID.String(),
			Resource: info.Resource,
			Method:   info.Method,
			Allowed:  info.Allowed,
			CreateAt: info.CreateAt,
		})
	}

	return &npool.GetAuthHistoriesResponse{
		Infos: ahs,
	}, nil
}

func GetByApp(ctx context.Context, in *npool.GetAuthHistoriesByAppRequest) (*npool.GetAuthHistoriesByAppResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, grpcTimeout)
	defer cancel()

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db client: %v", err)
	}

	appID, err := uuid.Parse(in.GetAppID())
	if err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	infos, err := cli.
		AuthHistory.
		Query().
		Where(
			authhistory.And(
				authhistory.AppID(appID),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query auth history: %v", err)
	}

	ahs := []*npool.AuthHistory{}
	for _, info := range infos {
		ahs = append(ahs, &npool.AuthHistory{
			ID:       info.ID.String(),
			AppID:    info.AppID.String(),
			UserID:   info.UserID.String(),
			Resource: info.Resource,
			Method:   info.Method,
			Allowed:  info.Allowed,
			CreateAt: info.CreateAt,
		})
	}

	return &npool.GetAuthHistoriesByAppResponse{
		Infos: ahs,
	}, nil
}
