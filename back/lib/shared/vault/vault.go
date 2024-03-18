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

type VaultAPI interface {
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
	HTTPAPI httpclient.HttpApi
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
	Secret Secret `json:"secret"`
}

type Secret struct {
	CreatedAt     time.Time `json:"created_at"`
	CreatedBy     CreatedBy `json:"created_by"`
	LatestVersion string    `json:"latest_version"`
	Name          string    `json:"name"`
	Version       Version   `json:"version"`
}
type CreatedBy struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Type  string `json:"type"`
}

type Version struct {
	Version   string    `json:"version"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	Value     string    `json:"value"`
	CreatedBy CreatedBy `json:"created_by"`
}

func InitializeVaultAPI(myHTTPAPI httpclient.HttpApi, vaultConfig Vault) (VaultClient, error) {
	vault, err := getAppVariables()
	if err != nil {
		return VaultClient{}, errors.Wrapf(err, "InitializeVaultAPI, unable to get app variables")
	}

	return VaultClient{Vault: Vault{
		AppName:        vaultConfig.AppName,
		AuthURL:        vaultConfig.AuthURL,
		BaseURL:        vaultConfig.BaseURL,
		OrganizationID: vault.OrganizationID,
		ProjectID:      vault.ProjectID,
	}, HTTPAPI: myHTTPAPI}, nil
}

func getClientVariables() (clientID, secret string, err error) {
	clientID = os.Getenv("HCP_CLIENT_ID")
	if clientID == "" {
		return "", "", errors.Wrapf(err, "getClientVariables, unable to get env variables")
	}

	secret = os.Getenv("HCP_CLIENT_SECRET")
	if secret == "" {
		return "", "", errors.Wrapf(err, "getClientVariables, unable to get env variables")
	}

	return clientID, secret, nil
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

func (vc *VaultClient) requestAccessKey() (VaultToken, error) {
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

	token, err := vc.getToken(fmt.Sprintf("%s/oauth/token", vc.Vault.AuthURL), headers, reqBody)
	if err != nil {
		return VaultToken{}, errors.Wrapf(err, "requestAccessKey, unable to get token")
	}

	return token, nil
}

func (vc *VaultClient) getToken(url string, headers map[string]string, reqBody TokenRequestBody) (VaultToken, error) {
	var token VaultToken

	jsonBody, err := httpclient.PrepareBody(reqBody)
	if err != nil {
		return VaultToken{}, errors.Wrapf(err, "getToken, unable to prepare body")
	}

	res, status, err := vc.HTTPAPI.Post(url, headers, jsonBody)
	if err != nil || status != http.StatusOK {
		return VaultToken{}, errors.Wrapf(err, "getToken, request impossible")
	}

	if err = json.Unmarshal(res, &token); err != nil {
		return VaultToken{}, errors.Wrapf(err, "getToken, unable to unmarshal")
	}

	return token, nil
}

func (vc *VaultClient) ReadSecret(secretName string) (string, error) {
	var appSecret AppSecret
	headers, err := vc.prepareReadSecretHeaders()
	if err != nil {
		return "", errors.Errorf("%s, ReadSecret, unable to get prepare headers to read secret", err)
	}

	url := fmt.Sprintf("%s/secrets/2023-06-13/organizations/%s/projects/%s/apps/%s/open/%s",
		vc.Vault.BaseURL, vc.Vault.OrganizationID, vc.Vault.ProjectID, vc.Vault.AppName, secretName)

	res, status, err := vc.HTTPAPI.Get(url, headers)
	if err != nil || status != http.StatusOK {
		return "", errors.Errorf("ReadSecret, request impossible %s", err)
	}

	return returnSecret(res, &appSecret)
}

func (vc *VaultClient) prepareReadSecretHeaders() (map[string]string, error) {
	token, err := vc.requestAccessKey()
	if err != nil {
		return nil, errors.Errorf("%s, ReadSecret, unable to get request access key", err)
	}

	headerLength := 2
	headers := make(map[string]string, headerLength)
	headers["content-type"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token.AccessToken)

	return headers, nil
}

func returnSecret(res []byte, appSecret *AppSecret) (string, error) {
	if err := json.Unmarshal(res, &appSecret); err != nil {
		return "", errors.Errorf("ReadSecret, unable to marshall %s", err)
	}

	return appSecret.Secret.Version.Value, nil
}
