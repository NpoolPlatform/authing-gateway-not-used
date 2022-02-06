package appauth

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/authinggateway"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
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

func CreateAppAuthForOtherApp(ctx context.Context, in *npool.CreateAppAuthForOtherAppRequest) (*npool.CreateAppAuthForOtherAppResponse, error) {
	return nil, nil
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
