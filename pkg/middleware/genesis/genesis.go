package genesis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	appusermgrconstant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	appusermgrpb "github.com/NpoolPlatform/message/npool/appuser/mgr/v1"
	npool "github.com/NpoolPlatform/message/npool/authinggateway"

	constant "github.com/NpoolPlatform/authing-gateway/pkg/const"
	appauthcrud "github.com/NpoolPlatform/authing-gateway/pkg/crud/appauth"
	appuserauthcrud "github.com/NpoolPlatform/authing-gateway/pkg/crud/appuserauth"
	grpc2 "github.com/NpoolPlatform/authing-gateway/pkg/grpc"
)

type genesisURL struct {
	Path   string
	Method string
}

func processGenesisURLs(urls []genesisURL, appID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	auths, err := appauthcrud.GetByApp(ctx, appID)
	if err != nil {
		logger.Sugar().Infof("fail get genesis app auth: %v", err)
		return
	}

	myURLs := []genesisURL{}
	for _, url := range urls {
		if url.Path == "" || url.Method == "" {
			logger.Sugar().Warnf("invalid url: %v", url)
			continue
		}

		found := false
		for _, info := range auths {
			if info.Resource == url.Path && info.Method == url.Method {
				found = true
				break
			}
		}
		if !found {
			myURLs = append(myURLs, url)
		}
	}

	for _, url := range myURLs {
		_, err := appauthcrud.CreateForOtherApp(ctx, &npool.CreateAppAuthForOtherAppRequest{
			Info: &npool.AppAuth{
				AppID:    appID,
				Resource: url.Path,
				Method:   url.Method,
			},
		})
		if err != nil {
			logger.Sugar().Errorf("fail create genesis url: %v", err)
			return
		}
	}
}

func watch() {
	hostname := config.GetStringValueWithNameSpace("", config.KeyHostname)
	urlsJSON := config.GetStringValueWithNameSpace(hostname, constant.KeyGenesisURLs)

	urls := []genesisURL{}
	err := json.Unmarshal([]byte(urlsJSON), &urls)
	if err == nil {
		logger.Sugar().Infof("process genesis urls: %v", urls)
		processGenesisURLs(urls, appusermgrconstant.GenesisAppID)
		processGenesisURLs(urls, appusermgrconstant.ChurchAppID)
	} else {
		logger.Sugar().Warnf("invalid urls %v: %v", urlsJSON, err)
	}
}

func Watch() {
	ticker := time.NewTicker(10 * time.Minute)
	for {
		watch()
		<-ticker.C
	}
}

func createAppUserAuths(ctx context.Context, appID string) ([]*npool.Auth, error) {
	hostname := config.GetStringValueWithNameSpace("", config.KeyHostname)
	apisJSON := config.GetStringValueWithNameSpace(hostname, constant.KeyGenesisAuthingAPIs)
	apis := []genesisURL{}
	err := json.Unmarshal([]byte(apisJSON), &apis)
	if err != nil {
		return nil, fmt.Errorf("fail parse genesis authing apis: %v", err)
	}
	if len(apis) == 0 {
		return nil, fmt.Errorf("genesis authing apis not available")
	}

	resp, err := grpc2.GetGenesisAppRoleUsersByOtherApp(ctx, &appusermgrpb.GetGenesisAppRoleUsersByOtherAppRequest{
		TargetAppID: appID,
	})
	if err != nil {
		return nil, fmt.Errorf("fail get genesis app role users: %v", err)
	}

	auths := []*npool.Auth{}

	for _, user := range resp.Infos {
		for _, api := range apis {
			resp1, err := appuserauthcrud.Create(ctx, &npool.CreateAppUserAuthRequest{
				Info: &npool.AppUserAuth{
					AppID:    appID,
					UserID:   user.UserID,
					Resource: api.Path,
					Method:   api.Method,
				},
			})
			if err != nil {
				return nil, fmt.Errorf("fail create app user auth: %v", err)
			}
			auths = append(auths, resp1.Info)
		}
	}

	return auths, nil
}

func CreateGenesisAppUserAuth(ctx context.Context, in *npool.CreateGenesisAppUserAuthRequest) (*npool.CreateGenesisAppUserAuthResponse, error) {
	allAuths := []*npool.Auth{}

	auths, err := createAppUserAuths(ctx, appusermgrconstant.GenesisAppID)
	if err != nil {
		return nil, fmt.Errorf("fail create genesis app user auths: %v", err)
	}
	allAuths = append(allAuths, auths...)

	auths, err = createAppUserAuths(ctx, appusermgrconstant.ChurchAppID)
	if err != nil {
		return nil, fmt.Errorf("fail create church app user auths: %v", err)
	}
	allAuths = append(allAuths, auths...)

	return &npool.CreateGenesisAppUserAuthResponse{
		Infos: allAuths,
	}, nil
}
