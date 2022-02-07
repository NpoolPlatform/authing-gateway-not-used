package genesis

import (
	"context"
	"encoding/json"
	"time"

	appusermgrconstant "github.com/NpoolPlatform/appuser-manager/pkg/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/authinggateway"

	constant "github.com/NpoolPlatform/authing-gateway/pkg/const"
	appauthcrud "github.com/NpoolPlatform/authing-gateway/pkg/crud/appauth"
)

type genesisURL struct {
	Path   string
	Method string
}

func processGenesisURLs(urls []genesisURL) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	auths, err := appauthcrud.GetByApp(ctx, appusermgrconstant.GenesisAppID)
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
				AppID:    appusermgrconstant.GenesisAppID,
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
		processGenesisURLs(urls)
	} else {
		logger.Sugar().Warnf("invalid urls %v: %v", urlsJSON, err)
	}

	// TODO: authorize authing apis to genesis user (need retry)
}

func Watch() {
	ticker := time.NewTicker(10 * time.Minute)
	for {
		watch()
		<-ticker.C
	}
}
