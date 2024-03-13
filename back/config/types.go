package config

import (
	c "github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/lib/shared/httpclient"
	"github.com/KrispyTech/airneis/lib/shared/vault"
)

var (
	productionConfig ProductionConfig
	stagingConfig    StagingConfig
)

var selectConfigProcessor = map[bool]func(Config) (Config, error){
	//Returns Production Config
	true: func(config Config) (Config, error) {
		return buildEnvironmentConfig(config, c.ProductionEnv, productionConfig)
	},

	//Returns Staging Config
	false: func(config Config) (Config, error) {
		return buildEnvironmentConfig(config, c.StagingEnv, stagingConfig)
	},
}

type Config struct {
	Env       Env `yaml:"env"`
	Handler   ClientHandler
	Processor any
}

type StagingConfig struct {
	Env struct {
		Name string `yaml:"name"`
	} `yaml:"staging"`
}

type ProductionConfig struct {
	Env struct {
		Name string `yaml:"name"`
	} `yaml:"production"`
}

type Env struct {
	BuildProduction bool    `yaml:"build_production"`
	Text            Text    `yaml:"text"`
	Secrets         Secrets `yaml:"secrets"`
}

type ClientHandler struct {
	httpClient  httpclient.HttpApi
	VaultClient vault.VaultClient
}

type HashicorpVault struct {
	AppName        string `yaml:"app_name"`
	AuthURL        string `yaml:"authentication-url"`
	BaseURL        string `yaml:"base_url"`
	OrganizationID string
	ProjectID      string
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
