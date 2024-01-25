package vault_test

import (
	"testing"

	"github.com/KrispyTech/airneis/lib/shared/httpclient"
	"github.com/KrispyTech/airneis/lib/shared/vault"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestVaultClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VaultClient")
}

var _ = Describe("ReadSecret", func() {
	Context("When we want to read a secret on HCP Vault Secrets", func() {
		payload := `{"secret":{"name":"hello_world_secret", "version":{"version":"1", "type":"kv", "created_at":"2024-01-15T21:48:33.292865Z", "value":"hello_world", "created_by":{"name":"john.doe@gmail.com", "type":"TYPE_USER", "email":"john.doe@gmail.com"}}, "created_at":"2024-01-15T21:48:33.292865Z", "latest_version":"1", "created_by":{"name":"john.doe@gmail.com", "type":"TYPE_USER", "email":"john.doe@gmail.com"}, "sync_status":{}}}`
		httpClient := &httpclient.HttpClientMock{Resp: []httpclient.MockResponse{
			{Data: []byte(payload), Status: 200}, {Data: []byte(payload), Status: 200},
		}}

		vlt, err := vault.InitializeVaultApi(httpClient, vault.Vault{
			AppName:        "Krispy",
			AuthURL:        "blabalbla",
			BaseURL:        "BASE",
			OrganizationID: "123",
			ProjectID:      "456",
		})

		It("should not return an error during initialization", func() {
			Expect(err).To(Not(HaveOccurred()))
		})

		secret, err := vlt.ReadSecret("hello_world_secret")
		It("should return the secret value without error", func() {
			Expect(err).To(Not(HaveOccurred()))
			Expect(secret).To(Equal("hello_world"))
		})
	})
})
