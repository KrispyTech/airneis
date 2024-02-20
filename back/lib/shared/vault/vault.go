package vault

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/KrispyTech/airneis/lib/shared/httpclient"
	"github.com/pkg/errors"
)

type VaultApi interface {
	ReadSecret(secretName string) (secret string, err error)
}

type Vault struct {
	AppName        string `yaml:"app_name"`
	AuthURL        string `yaml:"authentication-url"`
	BaseURL        string `yaml:"base_url"`
	OrganizationID string
	ProjectID      string
}

type VaultClient struct {
	HttpApi httpclient.HttpApi
	Vault   Vault
}

type VaultToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type TokenRequestBody struct {
	Audience     string `json:"audience"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
}

type AppSecret struct {
	Secret struct {
		CreatedAt time.Time `json:"created_at"`
		CreatedBy struct {
			Email string `json:"email"`
			Name  string `json:"name"`
			Type  string `json:"type"`
		} `json:"created_by"`
		LatestVersion string `json:"latest_version"`
		Name          string `json:"name"`
		Version       struct {
			Version   string    `json:"version"`
			Type      string    `json:"type"`
			CreatedAt time.Time `json:"created_at"`
			Value     string    `json:"value"`
			CreatedBy struct {
				Email string `json:"email"`
				Name  string `json:"name"`
				Type  string `json:"type"`
			} `json:"created_by"`
		} `json:"version"`
		SyncStatus struct {
		} `json:"sync_status"`
	} `json:"secret"`
}

func InitializeVaultApi(httpApi httpclient.HttpApi, vaultConfig Vault) (VaultClient, error) {
	vault, err := getAppVariables()
	if err != nil {
		return VaultClient{}, errors.Wrapf(err, "InitializeVaultApi, unable to get app variables")
	}

	return VaultClient{Vault: Vault{
		AppName:        vaultConfig.AppName,
		AuthURL:        vaultConfig.AuthURL,
		BaseURL:        vaultConfig.BaseURL,
		OrganizationID: vault.OrganizationID,
		ProjectID:      vault.ProjectID,
	}, HttpApi: httpApi}, nil
}

func getClientVariables() (clientID string, secret string, err error) {
	clientID = os.Getenv("HCP_CLIENT_ID")
	if clientID == "" {
		return "", "", errors.Wrapf(err, "getClientVariables, unable to get env variables")
	}

	secret = os.Getenv("HCP_CLIENT_SECRET")
	if secret == "" {
		return "", "", errors.Wrapf(err, "getClientVariables, unable to get env variables")
	}

	return
}

func getAppVariables() (vault Vault, err error) {
	organisationID := os.Getenv("HCP_ORG_ID")
	if organisationID == "" {
		return Vault{}, errors.Wrapf(err, "getAppVariables, unable to get env variables")
	}

	projectID := os.Getenv("HCP_PROJECT_ID")
	if projectID == "" {
		return Vault{}, errors.Wrapf(err, "getAppVariables, unable to get env variables")
	}

	return Vault{
		OrganizationID: organisationID,
		ProjectID:      projectID,
	}, nil
}

func (c *VaultClient) requestAccessKey() (token VaultToken, err error) {
	url := fmt.Sprintf("%s/oauth/token", c.Vault.AuthURL)
	headers := make(map[string]string)
	headers["content-type"] = "application/json"

	clientID, secret, err := getClientVariables()
	if err != nil {
		return VaultToken{}, errors.Wrapf(err, "requestAccessKey, unable to get client variables")
	}

	reqBody := TokenRequestBody{
		Audience:     "https://api.hashicorp.cloud",
		GrantType:    "client_credentials",
		ClientID:     clientID,
		ClientSecret: secret,
	}

	token, err = getToken(c, url, headers, reqBody)
	if err != nil {
		return VaultToken{}, errors.Wrapf(err, "requestAccessKey, unable to get token")
	}

	return
}

func getToken(vc *VaultClient, url string, headers map[string]string, reqBody TokenRequestBody) (token VaultToken, err error) {
	jsonBody, err := vc.HttpApi.PrepareBody(reqBody)
	if err != nil {
		return VaultToken{}, errors.Wrapf(err, "getToken, unable to prepare body")
	}

	res, status, err := vc.HttpApi.Post(url, headers, jsonBody)
	if err != nil || status != http.StatusOK {
		return VaultToken{}, errors.Wrapf(err, "getToken, request impossible")
	}

	if err = json.Unmarshal(res, &token); err != nil {
		return VaultToken{}, errors.Wrapf(err, "getToken, unable to unmarshal")
	}

	return
}

func (vc *VaultClient) ReadSecret(secretName string) (secret string, err error) {
	var appSecret AppSecret
	token, err := vc.requestAccessKey()
	if err != nil {
		return "", errors.Errorf("%s, ReadSecret, unable to get request access key", err)
	}

	headers := make(map[string]string, 2)
	headers["content-type"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token.AccessToken)

	url := fmt.Sprintf("%s/secrets/2023-06-13/organizations/%s/projects/%s/apps/%s/open/%s", vc.Vault.BaseURL, vc.Vault.OrganizationID, vc.Vault.ProjectID, vc.Vault.AppName, secretName)
	res, status, err := vc.HttpApi.Get(url, headers)
	if err != nil || status != http.StatusOK {
		return "", errors.Errorf(err, "ReadSecret, request impossible")
	}

	if err = json.Unmarshal(res, &appSecret); err != nil {
		return "", errors.Errorf(err, "ReadSecret, unable to unmarshal")
	}

	return appSecret.Secret.Version.Value, nil
}
