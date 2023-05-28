package checkers

import (
	"fmt"
	"github.com/taropowder/host-cdn-checker/config"
	"testing"
)

func TestCertificateChecker_Check(t *testing.T) {
	c := CertificateChecker{}
	config.Instance = &config.Config{}
	config.Instance.Certificates.BlackDomain = []string{"cdn"}
	fmt.Println(c.Check("124.236.18.244"))
}
