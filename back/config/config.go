package config

import (
	"airneis/lib/shared/httpclient"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

const stagingEnv string = "staging"
const productionEnv string = "production"

type Config struct {
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

	return config, nil
}
