package appuserauth

import (
	"context"
	"time"

	"github.com/NpoolPlatform/authing-gateway/pkg/db"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/appuserauth"
	npool "github.com/NpoolPlatform/message/npool/authinggateway"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

const (
	dbTimeout = 5 * time.Second
)

func validateAppUserAuth(info *npool.AppUserAuth) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return xerrors.Errorf("invalid role id: %v", err)
	}
	if info.GetResource() == "" || info.GetMethod() == "" {
		return xerrors.Errorf("invalid resource")
	}
	return nil
}

func dbRowToAuth(row *ent.AppUserAuth) *npool.Auth {
	return &npool.Auth{
		ID:       row.ID.String(),
		AppID:    row.AppID.String(),
		UserID:   row.UserID.String(),
		Resource: row.Resource,
		Method:   row.Method,
	}
}

func Create(ctx context.Context, in *npool.CreateAppUserAuthRequest) (*npool.CreateAppUserAuthResponse, error) {
	if err := validateAppUserAuth(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		AppUserAuth.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetResource(in.GetInfo().GetResource()).
		SetMethod(in.GetInfo().GetMethod()).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app user auth: %v", err)
	}

	return &npool.CreateAppUserAuthResponse{
		Info: dbRowToAuth(info),
	}, nil
}

func GetByAppUserResourceMethod(ctx context.Context, in *npool.GetAppUserAuthByAppUserResourceMethodRequest) (*npool.GetAppUserAuthByAppUserResourceMethodResponse, error) {
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
		AppUserAuth.
		Query().
		Where(
			appuserauth.And(
				appuserauth.AppID(uuid.MustParse(in.GetAppID())),
				appuserauth.UserID(uuid.MustParse(in.GetUserID())),
				appuserauth.Resource(in.GetResource()),
				appuserauth.Method(in.GetMethod()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app user auth: %v", err)
	}

	var appAuth *npool.Auth
	for _, info := range infos {
		appAuth = dbRowToAuth(info)
		break
	}

	return &npool.GetAppUserAuthByAppUserResourceMethodResponse{
		Info: appAuth,
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteAppUserAuthRequest) (*npool.DeleteAppUserAuthResponse, error) {
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
		AppUserAuth.
		UpdateOneID(id).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app user auth: %v", err)
	}

	return &npool.DeleteAppUserAuthResponse{
		Info: dbRowToAuth(info),
	}, nil
}
