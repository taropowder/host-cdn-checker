package checkers

import (
	"crypto/tls"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/taropowder/host-cdn-checker/config"
	"strings"
)

type CertificateChecker struct {
}

func (c *CertificateChecker) Check(ip string) (isCDN bool, trust bool, err error) {

	conf := config.Instance.Certificates

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:443", ip), &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Errorf("Error in Dial %v", err)
		return false, false, err
	}
	defer conn.Close()
	certs := conn.ConnectionState().PeerCertificates
	for _, certInfo := range certs {

		for _, domain := range conf.BlackDomain {
			if strings.Contains(certInfo.Subject.CommonName, domain) {
				log.Info(fmt.Sprintf("Subject.CommonName %s hit rule %s", certInfo.Subject.CommonName, domain))
				return true, true, nil
			}
		}

		for _, domain := range conf.WhiteDomain {
			if strings.Contains(certInfo.Subject.CommonName, domain) {
				return false, true, nil
			}
		}

	}

	//fmt.Println("过期时间:", certInfo.NotAfter)
	//fmt.Println("组织信息:", certInfo.Subject)
	//fmt.Println("颁发者:", certInfo.Issuer)
	//fmt.Println("CommonName:", certInfo.Subject.CommonName)
	//fmt.Println("Names:", certInfo.Subject.Names)
	//fmt.Println("DNS:", certInfo.DNSNames)

	//certInfo := resp.TLS.PeerCertificates[0]

	return false, false, nil
}
