package vault

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/KrispyTech/airneis/lib/shared/airror"

	"github.com/KrispyTech/airneis/lib/shared/httpclient"
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
	HTTPAPI httpclient.Req
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

func InitializeVaultAPI(myHTTPAPI httpclient.Req, vaultConfig Vault) (VaultClient, error) {
	vault, err := getAppVariables()
	if err != nil {
		return VaultClient{}, airror.New(err, "InitializeVaultAPI, unable to get app variables")
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
		return "", "", fmt.Errorf("getClientVariables, unable to get env variables")
	}

	secret = os.Getenv("HCP_CLIENT_SECRET")
	if secret == "" {
		return "", "", fmt.Errorf("getClientVariables, unable to get env variables")
	}

	return clientID, secret, nil
}

func getAppVariables() (vault Vault, err error) {
	organisationID := os.Getenv("HCP_ORG_ID")
	if organisationID == "" {
		return Vault{}, airror.New(err, "getAppVariables, unable to get : HCP_ORG_ID")
	}

	projectID := os.Getenv("HCP_PROJECT_ID")
	if projectID == "" {
		return Vault{}, airror.New(err, "getAppVariables, unable to get : HCP_PROJECT_ID")
	}

	return Vault{
		OrganizationID: organisationID,
		ProjectID:      projectID,
	}, nil
}

func (vc *VaultClient) requestAccessKey() (VaultToken, error) {
	headerLength := 1
	headers := make(map[string]string, headerLength)
	headers["content-type"] = "application/json"

	clientID, secret, err := getClientVariables()
	if err != nil {
		return VaultToken{}, airror.New(err, "requestAccessKey, unable to get client variables")
	}

	reqBody := TokenRequestBody{
		Audience:     "https://api.hashicorp.cloud",
		GrantType:    "client_credentials",
		ClientID:     clientID,
		ClientSecret: secret,
	}

	token, err := vc.getToken(fmt.Sprintf("%s/oauth2/token", vc.Vault.AuthURL), headers, reqBody)
	if err != nil {
		return VaultToken{}, airror.New(err, "requestAccessKey, unable to get token")
	}

	return token, nil
}

func (vc *VaultClient) getToken(url string, headers map[string]string, reqBody TokenRequestBody) (VaultToken, error) {
	var token VaultToken

	jsonBody, err := httpclient.PrepareBody(reqBody)
	if err != nil {
		return VaultToken{}, airror.New(err, "getToken, unable to prepare body")
	}

	res, status, err := vc.HTTPAPI.Post(url, headers, jsonBody)
	if err != nil || status != http.StatusOK {
		return VaultToken{}, airror.New(err, "getToken, request impossible")
	}

	if err = json.Unmarshal(res, &token); err != nil {
		return VaultToken{}, airror.New(err, "getToken, unable to unmarshal")
	}

	return token, nil
}

func (vc *VaultClient) ReadSecret(secretName string) (string, error) {
	var appSecret AppSecret
	headers, err := vc.prepareReadSecretHeaders()
	if err != nil {
		return "", airror.New(err, "ReadSecret, unable to get headers to read secret.")
	}

	url := fmt.Sprintf("%s/secrets/2023-06-13/organizations/%s/projects/%s/apps/%s/open/%s",
		vc.Vault.BaseURL, vc.Vault.OrganizationID, vc.Vault.ProjectID, vc.Vault.AppName, secretName)

	res, status, err := vc.HTTPAPI.Get(url, headers)
	if err != nil || status != http.StatusOK {
		return "", airror.New(err, "ReadSecret, request impossible")
	}

	return returnSecret(res, &appSecret)
}

func (vc *VaultClient) prepareReadSecretHeaders() (map[string]string, error) {
	token, err := vc.requestAccessKey()
	if err != nil {
		return nil, airror.New(err, "prepareReadSecretHeaders, unable to get request access key")
	}

	headerLength := 2
	headers := make(map[string]string, headerLength)
	headers["content-type"] = "application/json"
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token.AccessToken)

	return headers, nil
}

func returnSecret(res []byte, appSecret *AppSecret) (string, error) {
	if err := json.Unmarshal(res, &appSecret); err != nil {
		return "", airror.New(err, "returnSecret, unable to marshall")
	}

	return appSecret.Secret.Version.Value, nil
}
