package vault

import (
	"context"
	"encoding/json"
	"log"

	"github.com/hashicorp/vault-client-go"
	"github.com/pkg/errors"
)

type VaultAPI interface {
}

type VaultClient struct {
	Api *vault.Client
	ctx context.Context
}

func InitializeVaultClient() (VaultAPI, error) {
	client, err := vault.New()
	if err != nil {
		log.Fatal(err)
	}

	return &VaultClient{Api: client}, nil
}

func (c *VaultClient) ReadSecret(secretName string) (secret string, err error) {
	resp, err := c.Api.Secrets.KvV1Read(c.ctx, secretName)
	if err != nil {
		return "", errors.Wrapf(err, "ReadSecret, unable to read secret")
	}

	data, err := json.Marshal(resp.Data)
	if err != nil {
		return "", errors.Wrapf(err, "ReadSecret, unable to marshall")
	}

	return string(data), nil
}
