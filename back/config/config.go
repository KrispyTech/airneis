package config

import (
	"fmt"
	"os"

	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/lib/shared/httpclient"
	"github.com/KrispyTech/airneis/lib/shared/vault"
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

<<<<<<< HEAD
func InitializeConfig() (Config, error) {
	var config Config
=======
func InitializeConfig() (config Config, err error) {
>>>>>>> 0cb0e09 (Refactor the way we build our configuration)

	envFile, err := os.ReadFile("config/default.yaml")
	if err != nil {
		return Config{}, errors.Errorf("InitializeConfig, unable to read config file %s", err)
	}

	if err = yaml.Unmarshal(envFile, &config); err != nil {
		return Config{}, errors.Errorf("InitializeConfig, unable to unmarshall %s", err)
	}

	if err := godotenv.Load(); err != nil {
		return Config{}, errors.Errorf("Couldn't load .env file %s", err)
	}

	config.Handler, err = loadClientHandler(config)
	if err != nil {
		return Config{}, errors.Wrapf(err, "InitializeConfig, unable to load client handler")
	}

<<<<<<< HEAD
	initializedConfig, err := selectConfigProcessor[config.Env.BuildProduction](config)
=======
	config, err = selectConfigProcessor[config.Env.BuildProduction](config)
>>>>>>> 0cb0e09 (Refactor the way we build our configuration)
	if err != nil {
		return Config{}, errors.Wrapf(err, "buildEnvironmentConfig, unable to build")
	}

	return initializedConfig, nil
}

func buildEnvironmentConfig(config Config, env string, configProcessor any) (Config, error) {
	configFile, err := os.ReadFile(fmt.Sprintf("config/%s.yaml", env))
	if err != nil {
		return Config{}, errors.Errorf("loadConfig, unable to read config file %s", err)
	}

	if err = yaml.Unmarshal(configFile, &configProcessor); err != nil {
		return Config{}, errors.Errorf("loadConfig, unable to unmarshall config file %s", err)
	}

	config.Processor = configProcessor

	if err = initDatabase(config, env); err != nil {
		return Config{}, errors.Wrapf(err, "buildEnvironmentConfig, unable to initialize database")
	}

<<<<<<< HEAD
	log.Infof("%s environment has been loaded", helpers.Capitalize(env))
=======
	log.Info("Which Environment has been loaded")
>>>>>>> 0cb0e09 (Refactor the way we build our configuration)

	return config, nil
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
