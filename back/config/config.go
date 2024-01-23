package config

import (
	"airneis/lib/shared/httpclient"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Env     Env `yaml:"env"`
	Handler ClientHandler
}

type Env struct {
	Text    Text    `yaml:"text"`
	Secrets Secrets `yaml:"secrets"`
}

type ClientHandler struct {
	httpClient httpclient.HttpApi
}

type HashicorpVault struct {
	OrganizationID string `yaml:"organization_id"`
	ProjectID      string `yaml:"project_id"`
	AppName        string `yaml:"app_name"`
}

type Text struct {
	HashicorpVault HashicorpVault `yaml:"vault"`
}

type Secrets struct {
	StripeId      string `yaml:"STRIPE_ID"`
	StripeToken   string `yaml:"STRIPE_TOKEN"`
	SendgridId    string `yaml:"SENGDRID_ID"`
	SendgridToken string `yaml:"SENDGRID_TOKEN"`
}

func InitializeConfig() (config Config, err error) {
	envFile, err := os.ReadFile("config/default.yaml")
	if err != nil {
		return Config{}, errors.Wrapf(err, "InitializeConfig, unable to read config file")
	}

	if err = yaml.Unmarshal(envFile, &config); err != nil {
		return Config{}, errors.Wrapf(err, "InitializeConfig, unable to unmarshal")
	}

	config.Handler = ClientHandler{httpClient: httpclient.InitializeHttpApi()}
	return
}
