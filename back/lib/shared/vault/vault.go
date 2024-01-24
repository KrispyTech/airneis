package vault

import (
	"airneis/lib/shared/httpclient"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

type VaultApi interface {
	ReadSecret(secretName string) (secret string, err error)
}

type Vault struct {
	BaseURL        string `yaml:"base_url"`
	AuthURL        string `yaml:"authentication-url"`
	OrganizationID string `yaml:"organization_id"`
	ProjectID      string `yaml:"project_id"`
	AppName        string `yaml:"app_name"`
}

type VaultClient struct {
	Vault      Vault
	HttpClient httpclient.HttpClient
}

type VaultToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

type TokenRequestBody struct {
	Audience     string `json:"audience"`
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func CreateVaultClient(httpClient httpclient.HttpClient, vaultConfig Vault) (VaultClient, error) {
	return VaultClient{Vault: vaultConfig, HttpClient: httpClient}, nil
}

func InitializeVaultApi(vaultClient VaultClient) (VaultApi, error) {
	return &vaultClient, nil
}

func getClientVariables() (clientID string, secret string, err error) {
	clientID = os.Getenv("HCP_CLIENT_ID")
	if clientID == "" {
		return "", "", errors.Wrapf(err, "generateAccessKey, unable to get env variables")
	}

	secret = os.Getenv("HCP_CLIENT_SECRET")
	if secret == "" {
		return "", "", errors.Wrapf(err, "generateAccessKey, unable to get env variables")
	}

	return
}

func (c *VaultClient) requestAccessKey() (token VaultToken, err error) {
	url := fmt.Sprintf("%s/oauth2/token", c.Vault.AuthURL)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	clientID, secret, err := getClientVariables()
	if err != nil {
		return VaultToken{}, err
	}

	reqBody := TokenRequestBody{
		Audience:     "https://api.hashicorp.cloud",
		GrantType:    "client_credentials",
		ClientID:     clientID,
		ClientSecret: secret,
	}

	token, err = getToken(c, url, headers, reqBody)
	if err != nil {
		return VaultToken{}, errors.Wrapf(err, "generateAccessKey, unable to get token")
	}

	return
}

func getToken(vc *VaultClient, url string, headers map[string]string, reqBody TokenRequestBody) (token VaultToken, err error) {
	jsonBody, err := vc.HttpClient.PrepareBody(reqBody)
	if err != nil {
		return VaultToken{}, errors.Wrapf(err, "getToken, unable to prepare body")
	}

	res, status, err := vc.HttpClient.Post(url, headers, jsonBody)
	if err != nil || status != http.StatusOK {
		return VaultToken{}, errors.Wrapf(err, "getToken, request impossible")
	}

	if err = json.Unmarshal(res, &token); err != nil {
		return VaultToken{}, errors.Wrapf(err, "getToken, unable to unmarshal")
	}

	return
}

func (c *VaultClient) ReadSecret(secretName string) (secret string, err error) {
	token, err := c.requestAccessKey()
	if err != nil {
		return "", errors.Wrapf(err, "ReadSecret, unable to get request access key")
	}

	fmt.Println(token.ExpiresIn)
	return "", nil
}
