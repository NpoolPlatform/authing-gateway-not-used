package approleauth

import (
	"context"
	"time"

	"github.com/NpoolPlatform/authing-gateway/pkg/db"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/approleauth"
	npool "github.com/NpoolPlatform/message/npool/authinggateway"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

const (
	dbTimeout = 5 * time.Second
)

func validateAppRoleAuth(info *npool.AppRoleAuth) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetRoleID()); err != nil {
		return xerrors.Errorf("invalid role id: %v", err)
	}
	if info.GetResource() == "" || info.GetMethod() == "" {
		return xerrors.Errorf("invalid resource")
	}
	return nil
}

func dbRowToAuth(row *ent.AppRoleAuth) *npool.Auth {
	return &npool.Auth{
		ID:       row.ID.String(),
		AppID:    row.AppID.String(),
		RoleID:   row.RoleID.String(),
		Resource: row.Resource,
		Method:   row.Method,
	}
}

func Create(ctx context.Context, in *npool.CreateAppRoleAuthRequest) (*npool.CreateAppRoleAuthResponse, error) {
	if err := validateAppRoleAuth(in.GetInfo()); err != nil {
		return nil, xerrors.Errorf("invalid parameter: %v", err)
	}

	cli, err := db.Client()
	if err != nil {
		return nil, xerrors.Errorf("fail get db: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	err = cli.
		AppRoleAuth.
		Create().
		SetAppID(uuid.MustParse(in.GetInfo().GetAppID())).
		SetRoleID(uuid.MustParse(in.GetInfo().GetRoleID())).
		SetResource(in.GetInfo().GetResource()).
		SetMethod(in.GetInfo().GetMethod()).
		OnConflict().
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail create app role auth: %v", err)
	}

	resp, err := GetByAppRoleResourceMethod(ctx, &npool.GetAppRoleAuthByAppRoleResourceMethodRequest{
		AppID:    in.GetInfo().GetAppID(),
		RoleID:   in.GetInfo().GetRoleID(),
		Resource: in.GetInfo().GetResource(),
		Method:   in.GetInfo().GetMethod(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app role auth: %v", err)
	}

	return &npool.CreateAppRoleAuthResponse{
		Info: resp.Info,
	}, nil
}

func GetByAppRoleResourceMethod(ctx context.Context, in *npool.GetAppRoleAuthByAppRoleResourceMethodRequest) (*npool.GetAppRoleAuthByAppRoleResourceMethodResponse, error) {
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
		AppRoleAuth.
		Query().
		Where(
			approleauth.And(
				approleauth.AppID(uuid.MustParse(in.GetAppID())),
				approleauth.RoleID(uuid.MustParse(in.GetRoleID())),
				approleauth.Resource(in.GetResource()),
				approleauth.Method(in.GetMethod()),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app role auth: %v", err)
	}

	var appAuth *npool.Auth
	for _, info := range infos {
		appAuth = dbRowToAuth(info)
		break
	}

	return &npool.GetAppRoleAuthByAppRoleResourceMethodResponse{
		Info: appAuth,
	}, nil
}

func Delete(ctx context.Context, in *npool.DeleteAppRoleAuthRequest) (*npool.DeleteAppRoleAuthResponse, error) {
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
		AppRoleAuth.
		UpdateOneID(id).
		SetDeleteAt(uint32(time.Now().Unix())).
		Save(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail update app role auth: %v", err)
	}

	return &npool.DeleteAppRoleAuthResponse{
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
		AppRoleAuth.
		Query().
		Where(
			approleauth.AppID(uuid.MustParse(appID)),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app role auth: %v", err)
	}

	appAuths := []*npool.Auth{}
	for _, info := range infos {
		appAuths = append(appAuths, dbRowToAuth(info))
	}

	return appAuths, nil
}

func GetByAppResourceMethod(ctx context.Context, appID, resource, method string) ([]*npool.Auth, error) {
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
		AppRoleAuth.
		Query().
		Where(
			approleauth.And(
				approleauth.AppID(uuid.MustParse(appID)),
				approleauth.Resource(resource),
				approleauth.Method(method),
			),
		).
		All(ctx)
	if err != nil {
		return nil, xerrors.Errorf("fail query app role uth: %v", err)
	}

	appAuths := []*npool.Auth{}
	for _, info := range infos {
		appAuths = append(appAuths, dbRowToAuth(info))
	}

	return appAuths, nil
}
