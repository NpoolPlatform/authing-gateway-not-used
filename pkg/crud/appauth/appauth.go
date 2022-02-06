package appauth

import (
	"context"
	"time"

	"github.com/NpoolPlatform/authing-gateway/pkg/db"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/authinggateway"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

const (
	dbTimeout = 5 * time.Second
)

func validateAppAuth(info *npool.AppAuth) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if info.GetResource() == "" || info.GetMethod() == "" {
		return xerrors.Errorf("invalid resource")
	}
	return nil
}

func dbRowToAppAuth(row *ent.AppAuth) *npool.AppAuth {
	return &npool.AppAuth{
		ID:       row.ID.String(),
		AppID:    row.AppID.String(),
		Resource: row.Resource,
		Method:   row.Method,
	}
}

func CreateForOtherApp(ctx context.Context, in *npool.CreateAppAuthForOtherAppRequest) (*npool.CreateAppAuthForOtherAppResponse, error) {
	if err := validateAppAuth(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		AppAuth.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetResource(in.GetInfo().GetResource()).
		SetMethod(in.GetInfo().GetMethod()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app auth: %v", err)
	}

	return &npool.CreateAppAuthForOtherAppResponse{
		Info: dbRowToAppAuth(info),
	}, nil
}

func GetAppAuthByAppResourceMethod(ctx context.Context, in *npool.GetAppAuthByAppResourceMethodRequest) (*npool.GetAppAuthByAppResourceMethodResponse, error) {
	return nil, nil
}

func GetAppAuthByOtherAppResourceMethod(ctx context.Context, in *npool.GetAppAuthByOtherAppResourceMethodRequest) (*npool.GetAppAuthByOtherAppResourceMethodResponse, error) {
	return nil, nil
}

func DeleteAppAuth(ctx context.Context, in *npool.DeleteAppAuthRequest) (*npool.DeleteAppAuthResponse, error) {
	return nil, nil
}
