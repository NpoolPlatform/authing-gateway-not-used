package appuserauth

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/authinggateway"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func validateAppUserAuth(info *npool.AppUserAuth) error {
	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		return xerrors.Errorf("invalid app id: %v", err)
	}
	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		return xerrors.Errorf("invalid user id: %v", err)
	}
	if info.GetResource() == "" || info.GetMethod() == "" {
		return xerrors.Errorf("invalid resource")
	}
	return nil
}

func Create(ctx context.Context, in *npool.CreateAppUserAuthRequest) (*npool.CreateAppUserAuthResponse, error) {
	return nil, nil
}

func CreateForOtherApp(ctx context.Context, in *npool.CreateAppUserAuthForOtherAppRequest) (*npool.CreateAppUserAuthForOtherAppResponse, error) {
	return nil, nil
}

func GetByAppUserResourceMethod(ctx context.Context, in *npool.GetAppUserAuthByAppUserResourceMethodRequest) (*npool.GetAppUserAuthByAppUserResourceMethodResponse, error) {
	return nil, nil
}

func GetByOtherAppUserResourceMethod(ctx context.Context, in *npool.GetAppUserAuthByOtherAppUserResourceMethodRequest) (*npool.GetAppUserAuthByOtherAppUserResourceMethodResponse, error) {
	return nil, nil
}

func DeleteAppUserAuth(ctx context.Context, in *npool.DeleteAppUserAuthRequest) (*npool.DeleteAppUserAuthResponse, error) {
	return nil, nil
}
