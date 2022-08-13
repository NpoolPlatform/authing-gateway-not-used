package appuserauth

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/authing-gateway/pkg/db"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/appuserauth"
	npool "github.com/NpoolPlatform/message/npool/authinggateway"

	"github.com/google/uuid"
)

const (
	dbTimeout = 5 * time.Second
)

func validateAppUserAuth(info *npool.AppUserAuth) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return fmt.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return fmt.Errorf("invalid user id: %v", err)
	}
	if info.GetResource() == "" || info.GetMethod() == "" {
		return fmt.Errorf("invalid resource")
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
		return nil, fmt.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	err = cli.
		AppUserAuth.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetUserID(uuid.MustParse(in.GetInfo().GetUserID())).
		SetResource(in.GetInfo().GetResource()).
		SetMethod(in.GetInfo().GetMethod()).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail create app user auth: %v", err)
	}

	resp, err := GetByAppUserResourceMethod(ctx, &npool.GetAppUserAuthByAppUserResourceMethodRequest{
		AppID:    in.GetInfo().GetAppID(),
		UserID:   in.GetInfo().GetUserID(),
		Resource: in.GetInfo().GetResource(),
		Method:   in.GetInfo().GetMethod(),
	})
	if err != nil {
		return nil, fmt.Errorf("fail get app user auth: %v", err)
	}

	return &npool.CreateAppUserAuthResponse{
		Info: resp.Info,
	}, nil
}

func GetByAppUserResourceMethod(ctx context.Context, in *npool.GetAppUserAuthByAppUserResourceMethodRequest) (*npool.GetAppUserAuthByAppUserResourceMethodResponse, error) {
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db: %v", err)
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
		return nil, fmt.Errorf("fail query app user auth: %v", err)
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
		return nil, fmt.Errorf("invalid id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	info, err := cli.
		AppUserAuth.
		UpdateOneID(id).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail update app user auth: %v", err)
	}

	return &npool.DeleteAppUserAuthResponse{
		Info: dbRowToAuth(info),
	}, nil
}

func GetByApp(ctx context.Context, appID string) ([]*npool.Auth, error) {
	if _, err := uuid.Parse(appID); err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		AppUserAuth.
		Query().
		Where(
			appuserauth.AppID(uuid.MustParse(appID)),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app user auth: %v", err)
	}

	appAuths := []*npool.Auth{}
	for _, info := range infos {
		appAuths = append(appAuths, dbRowToAuth(info))
	}

	return appAuths, nil
}

func GetByAppResourceMethod(ctx context.Context, appID, resource, method string) ([]*npool.Auth, error) {
	if _, err := uuid.Parse(appID); err != nil {
		return nil, fmt.Errorf("invalid app id: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, fmt.Errorf("fail get db: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	infos, err := cli.
		AppUserAuth.
		Query().
		Where(
			appuserauth.And(
				appuserauth.AppID(uuid.MustParse(appID)),
				appuserauth.Resource(resource),
				appuserauth.Method(method),
			),
		).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail query app user uth: %v", err)
	}

	appAuths := []*npool.Auth{}
	for _, info := range infos {
		appAuths = append(appAuths, dbRowToAuth(info))
	}

	return appAuths, nil
}
