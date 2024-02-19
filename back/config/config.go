package config

import (
	"fmt"
	"os"

	neon "github.com/KrispyTech/airneis/db"
	"github.com/KrispyTech/airneis/lib/shared/httpclient"
	"github.com/KrispyTech/airneis/lib/shared/vault"

	"gorm.io/gorm"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const (
	stagingEnv    string = "staging"
	productionEnv string = "production"
)

type Config struct {
	DB         *gorm.DB
	Neon       neon.Database
	Env        Env `yaml:"env"`
	Handler    ClientHandler
	Production ProductionConfig
	Staging    StagingConfig
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

func InitializeConfig() (config Config, err error) {
	envFile, err := os.ReadFile("config/default.yaml")
	if err != nil {
		return Config{}, errors.Wrapf(err, "InitializeConfig, unable to read config file")
	}

	if err = yaml.Unmarshal(envFile, &config); err != nil {
		return Config{}, errors.Wrapf(err, "InitializeConfig, unable to unmarshal")
	}

	if err := godotenv.Load(); err != nil {
		return Config{}, errors.Wrapf(err, "Couldn't load .env file")
	}

	config.Handler, err = loadClientHandler(config)
	if err != nil {
		return Config{}, errors.Wrapf(err, "InitializeConfig, unable to load client handler")
	}

	if !config.Env.BuildProduction {
		config, err = loadConfig(config, stagingEnv)
		if err != nil {
			return Config{}, errors.Wrapf(err, "InitializeConfig, unable to load staging config")
		}
		logrus.Info("Staging Environment has been loaded")
	} else {
		config, err = loadConfig(config, productionEnv)
		if err != nil {
			return Config{}, errors.Wrapf(err, "InitializeConfig, unable to load config")
		}
		logrus.Info("Production Environment has been loaded")
	}

	return
}

func loadClientHandler(config Config) (ClientHandler, error) {
	httpClient := httpclient.InitializeHttpApi()
	vc, err := vault.InitializeVaultApi(httpClient, vault.Vault(config.Env.Text.HashicorpVault))
	if err != nil {
		return ClientHandler{}, errors.Wrapf(err, "loadClientHandler, unable to create vault client")
	}

	ch := ClientHandler{
		httpClient:  httpClient,
		VaultClient: vc,
	}
	return ch, nil
}

func loadConfig(config Config, envName string) (Config, error) {
	var stagingConfig StagingConfig
	var productionConfig ProductionConfig

	configFile, err := os.ReadFile(fmt.Sprintf("config/%s.yaml", envName))
	if err != nil {
		return Config{}, errors.Wrapf(err, "loadConfig, unable to read config file")
	}

	switch envName {
	case stagingEnv:
		if err = yaml.Unmarshal(configFile, &stagingConfig); err != nil {
			return Config{}, errors.Wrapf(err, "loadConfig, unable to unmarshal")
		}
		config.Staging = stagingConfig

	default:
		if err = yaml.Unmarshal(configFile, &productionConfig); err != nil {
			return Config{}, errors.Wrapf(err, "loadConfig, unable to unmarshal")
		}
		config.Production = productionConfig
	}

	config, err = InitDatabase(config)

	if err != nil {
		return Config{}, errors.Wrapf(err, "loadConfig, Unable to Init database")
	}

	return config, nil
}
