package authing

import (
	"context"

	appauthcrud "github.com/NpoolPlatform/authing-gateway/pkg/crud/appauth"
	approleauthcrud "github.com/NpoolPlatform/authing-gateway/pkg/crud/approleauth"
	appuserauthcrud "github.com/NpoolPlatform/authing-gateway/pkg/crud/appuserauth"
	authhistorycrud "github.com/NpoolPlatform/authing-gateway/pkg/crud/authhistory"
	grpc2 "github.com/NpoolPlatform/authing-gateway/pkg/grpc"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	appusermgrpb "github.com/NpoolPlatform/message/npool/appusermgr"
	npool "github.com/NpoolPlatform/message/npool/authinggateway"
	logingwpb "github.com/NpoolPlatform/message/npool/logingateway"

	"github.com/google/uuid"
	"golang.org/x/xerrors"
)

func AuthByApp(ctx context.Context, in *npool.AuthByAppRequest) (*npool.AuthByAppResponse, error) {
	allowed := false

	defer func() {
		err := authhistorycrud.Create(ctx, &npool.AuthHistory{
			AppID:    in.GetAppID(),
			UserID:   uuid.UUID{}.String(),
			Resource: in.GetResource(),
			Method:   in.GetMethod(),
			Allowed:  allowed,
		})
		if err != nil {
			logger.Sugar().Errorf("fail create auth history: %v", err)
		}
	}()

	resp, err := grpc2.GetAppInfo(ctx, &appusermgrpb.GetAppInfoRequest{
		ID: in.GetAppID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app: %v", err)
	}
	if resp.Info == nil {
		return nil, xerrors.Errorf("fail get app")
	}

	allowed = resp.Info.Ban == nil
	if !allowed {
		return &npool.AuthByAppResponse{
			Allowed: false,
		}, nil
	}

	resp1, err := appauthcrud.GetByAppResourceMethod(ctx, &npool.GetAppAuthByAppResourceMethodRequest{
		AppID:    in.GetAppID(),
		Resource: in.GetResource(),
		Method:   in.GetMethod(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app auth by resource method: %v", err)
	}
	if resp1.Info != nil {
		allowed = true
		return &npool.AuthByAppResponse{
			Allowed: allowed,
		}, nil
	}

	auths, err := approleauthcrud.GetByAppResourceMethod(ctx, in.GetAppID(), in.GetResource(), in.GetMethod())
	if err != nil {
		return nil, xerrors.Errorf("fail get app role auth by app resource method: %v", err)
	}
	if len(auths) > 0 {
		return &npool.AuthByAppResponse{
			Allowed: allowed,
		}, nil
	}

	auths, err = appuserauthcrud.GetByAppResourceMethod(ctx, in.GetAppID(), in.GetResource(), in.GetMethod())
	if err != nil {
		return nil, xerrors.Errorf("fail get app userg auth by app resource method: %v", err)
	}
	if len(auths) > 0 {
		allowed = true
	}

	return &npool.AuthByAppResponse{
		Allowed: allowed,
	}, nil
}

func AuthByAppRoleUser(ctx context.Context, in *npool.AuthByAppRoleUserRequest) (*npool.AuthByAppRoleUserResponse, error) {
	allowed := false

	defer func() {
		err := authhistorycrud.Create(ctx, &npool.AuthHistory{
			AppID:    in.GetAppID(),
			UserID:   in.GetUserID(),
			Resource: in.GetResource(),
			Method:   in.GetMethod(),
			Allowed:  allowed,
		})
		if err != nil {
			logger.Sugar().Errorf("fail create auth history: %v", err)
		}
	}()

	resp, err := AuthByApp(ctx, &npool.AuthByAppRequest{
		AppID: in.GetAppID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail auth by app: %v", err)
	}
	if !resp.Allowed {
		return &npool.AuthByAppRoleUserResponse{
			Allowed: false,
		}, nil
	}

	resp1, err := grpc2.GetAppUserInfo(ctx, &appusermgrpb.GetAppUserInfoRequest{
		ID: in.GetUserID(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app user: %v", err)
	}
	if resp1.Info == nil {
		return &npool.AuthByAppRoleUserResponse{
			Allowed: false,
		}, nil
	}

	if resp1.Info.Ban != nil {
		return &npool.AuthByAppRoleUserResponse{
			Allowed: false,
		}, nil
	}

	resp2, err := grpc2.Logined(ctx, &logingwpb.LoginedRequest{
		AppID:  in.GetAppID(),
		UserID: in.GetUserID(),
		Token:  in.GetToken(),
	})
	if err != nil {
		return nil, xerrors.Errorf("user not login: %v", err)
	}
	if resp2.Info == nil {
		return nil, xerrors.Errorf("user not login")
	}
	if resp2.Info.Roles == nil {
		return nil, xerrors.Errorf("invalid user roles")
	}

	resp3, err := appuserauthcrud.GetByAppUserResourceMethod(ctx, &npool.GetAppUserAuthByAppUserResourceMethodRequest{
		AppID:    in.GetAppID(),
		UserID:   in.GetUserID(),
		Resource: in.GetResource(),
		Method:   in.GetMethod(),
	})
	if err != nil {
		return nil, xerrors.Errorf("fail get app user auth by app user resource method: %v", err)
	}
	if resp3.Info != nil {
		allowed = true
		return &npool.AuthByAppRoleUserResponse{
			Allowed: allowed,
		}, nil
	}

	for _, role := range resp2.Info.Roles {
		resp, err := approleauthcrud.GetByAppRoleResourceMethod(ctx, &npool.GetAppRoleAuthByAppRoleResourceMethodRequest{
			AppID:    in.GetAppID(),
			RoleID:   role.ID,
			Resource: in.GetResource(),
			Method:   in.GetMethod(),
		})
		if err != nil {
			return nil, xerrors.Errorf("fail get app role auth by app role resource method: %v", err)
		}
		if resp.Info != nil {
			allowed = true
			return &npool.AuthByAppRoleUserResponse{
				Allowed: allowed,
			}, nil
		}
	}

	return &npool.AuthByAppRoleUserResponse{
		Allowed: false,
	}, nil
}

func GetAuthsByAppRole(ctx context.Context, in *npool.GetAuthsByAppRoleRequest) (*npool.GetAuthsByAppRoleResponse, error) {
	auths, err := approleauthcrud.GetByApp(ctx, in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("fail get by app: %v", err)
	}
	return &npool.GetAuthsByAppRoleResponse{
		Infos: auths,
	}, nil
}

func GetAuthsByAppUser(ctx context.Context, in *npool.GetAuthsByAppUserRequest) (*npool.GetAuthsByAppUserResponse, error) {
	auths, err := appuserauthcrud.GetByApp(ctx, in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("fail get by app: %v", err)
	}
	return &npool.GetAuthsByAppUserResponse{
		Infos: auths,
	}, nil
}

func GetAuthsByApp(ctx context.Context, in *npool.GetAuthsByAppRequest) (*npool.GetAuthsByAppResponse, error) {
	appAuths := []*npool.Auth{}

	auths, err := approleauthcrud.GetByApp(ctx, in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("fail get by app: %v", err)
	}
	appAuths = append(appAuths, auths...)

	auths, err = appuserauthcrud.GetByApp(ctx, in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("fail get by app: %v", err)
	}
	appAuths = append(appAuths, auths...)

	auths, err = appauthcrud.GetByApp(ctx, in.GetAppID())
	if err != nil {
		return nil, xerrors.Errorf("fail get by app: %v", err)
	}
	appAuths = append(appAuths, auths...)

	return &npool.GetAuthsByAppResponse{
		Infos: appAuths,
	}, nil
}
