package checkers

import (
	"crypto/tls"
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

func Test2(t *testing.T) {
	//conn, err := tls.Dial("tcp", "103.254.188.41:443", nil)
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "103.254.188.41:443", conf)
	if err != nil {
		fmt.Println("Error in Dial", err)
		return
	}
	defer conn.Close()
	certs := conn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		fmt.Printf("Issuer Name: %s\n", cert.Issuer)
		fmt.Printf("Expiry: %s \n", cert.NotAfter.Format("2006-January-02"))
		fmt.Printf("Common Name: %s \n", cert.Issuer.CommonName)

	}
}
