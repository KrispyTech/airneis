package vault_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestVaultClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VaultClient")
}
