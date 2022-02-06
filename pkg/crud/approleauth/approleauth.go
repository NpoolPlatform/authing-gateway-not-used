package approleauth

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/authinggateway"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
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

func CreateAppRoleAuth(ctx context.Context, in *npool.CreateAppRoleAuthRequest) (*npool.CreateAppRoleAuthResponse, error) {
	return nil, nil
}

func CreateAppRoleAuthForOtherApp(ctx context.Context, in *npool.CreateAppRoleAuthForOtherAppRequest) (*npool.CreateAppRoleAuthForOtherAppResponse, error) {
	return nil, nil
}

func GetAppAuthByAppRoleResourceMethod(ctx context.Context, in *npool.GetAppAuthByAppRoleResourceMethodRequest) (*npool.GetAppAuthByAppRoleResourceMethodResponse, error) {
	return nil, nil
}

func GetAppAuthByOtherAppRoleResourceMethod(ctx context.Context, in *npool.GetAppAuthByOtherAppRoleResourceMethodRequest) (*npool.GetAppAuthByOtherAppRoleResourceMethodResponse, error) {
	return nil, nil
}

func DeleteAppRoleAuth(ctx context.Context, in *npool.DeleteAppRoleAuthRequest) (*npool.DeleteAppRoleAuthResponse, error) {
	return nil, nil
}
