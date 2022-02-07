package appauth

import (
	"context"
	"time"

	"github.com/NpoolPlatform/authing-gateway/pkg/db"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/appauth"
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

func dbRowToAuth(row *ent.AppAuth) *npool.Auth {
	return &npool.Auth{
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

	err = cli.
		AppAuth.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetResource(in.GetInfo().GetResource()).
		SetMethod(in.GetInfo().GetMethod()).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app auth: %v", err)
	}

	resp, err := GetByAppResourceMethod(ctx, &npool.GetAppAuthByAppResourceMethodRequest{
		AppID:    in.GetInfo().GetAppID(),
		Resource: in.GetInfo().GetResource(),
		Method:   in.GetInfo().GetMethod(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app auth: %v", err)
	}
	if resp.Info == nil {
		return nil, xerrors.Errorf("fail get app auth")
	}

	return &npool.CreateAppAuthForOtherAppResponse{
		Info: resp.Info,
	}, nil
}

func GetByAppResourceMethod(ctx context.Context, in *npool.GetAppAuthByAppResourceMethodRequest) (*npool.GetAppAuthByAppResourceMethodResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		AppAuth.
		Query().
		Where(
			appauth.And(
				appauth.AppID(uuid.MustParse(in.GetAppID())),
				appauth.Resource(in.GetResource()),
				appauth.Method(in.GetMethod()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app auth: %v", err)
	}

	var appAuth *npool.Auth
	for _, info := range infos {
		appAuth = dbRowToAuth(info)
		break
	}

	return &npool.GetAppAuthByAppResourceMethodResponse{
		Info: appAuth,
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteAppAuthRequest) (*npool.DeleteAppAuthResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return nil, xerrors.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		AppAuth.
		UpdateOneID(id).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app auth: %v", err)
	}

	return &npool.DeleteAppAuthResponse{
		Info: dbRowToAuth(info),
	}, nil
}

func GetByApp(ctx context.Context, appID string) ([]*npool.Auth, error) {
	if _, err := uuid.Parse(appID); err != nil {
		return nil, xerrors.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		AppAuth.
		Query().
		Where(
			appauth.AppID(uuid.MustParse(appID)),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app auth: %v", err)
	}

	appAuths := []*npool.Auth{}
	for _, info := range infos {
		appAuths = append(appAuths, dbRowToAuth(info))
	}

	return appAuths, nil
}
