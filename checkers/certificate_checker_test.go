package checkers

import (
	"fmt"
	"github.com/taropowder/host-cdn-checker/config"
	"testing"
)

func TestCertificateChecker_Check(t *testing.T) {
	c := CertificateChecker{}
	config.Instance = &config.Config{}
	config.Instance.Certificates.BlackDomain = []string{"cdn", "chinanetcenter.com"}
	fmt.Println(c.Check("103.254.188.41"))
}
