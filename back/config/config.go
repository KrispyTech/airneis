package config

import (
	"fmt"
	"os"

	"github.com/KrispyTech/airneis/lib/shared/airror"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/lib/shared/httpclient"
	"github.com/KrispyTech/airneis/lib/shared/vault"
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

func InitializeConfig() (Config, error) {
	var config Config

	envFile, err := os.ReadFile("config/default.yaml")
	if err != nil {
		return Config{}, airror.ReadError("os.ReadFile", err)
	}

	if err = yaml.Unmarshal(envFile, &config); err != nil {
		return Config{}, airror.UnmarshallError("yaml.UnMarshal", err)
	}

	if err := godotenv.Load(); err != nil {
		return Config{}, airror.LoadError("godotenv.Load", err)
	}

	config.Handler, err = loadClientHandler(config)
	if err != nil {
		return Config{}, airror.WrapLoadError("loadClientHandler", err)
	}

	initializedConfig, err := selectConfigProcessor[config.Env.BuildProduction](config)
	if err != nil {
		return Config{}, airror.WrapBuildError("buildEnvironmentConfig", err)
	}

	return initializedConfig, nil
}

func buildEnvironmentConfig(config Config, env string, configProcessor any) (Config, error) {
	configFile, err := os.ReadFile(fmt.Sprintf("config/%s.yaml", env))
	if err != nil {
		return Config{}, airror.ReadError("os.ReadFile", err)
	}

	if err = yaml.Unmarshal(configFile, &configProcessor); err != nil {
		return Config{}, airror.UnmarshallError("yaml.Unmarshal", err)
	}

	config.Processor = configProcessor

	if err = initializeDatabase(config, env); err != nil {
		return Config{}, airror.WrapBuildError("initializeDatabase", err)
	}

	log.Infof("%s environment has loaded", helpers.Capitalize(env))

	return config, nil
}

func loadClientHandler(config Config) (ClientHandler, error) {
	httpClient := httpclient.InitializeHttpApi()
	vc, err := vault.InitializeVaultApi(httpClient, vault.Vault(config.Env.Text.HashicorpVault))
	if err != nil {
		return ClientHandler{}, airror.WrapLoadError("loadClientHandler", err)
	}

	return ClientHandler{
		httpClient:  httpClient,
		VaultClient: vc,
	}, nil
}
