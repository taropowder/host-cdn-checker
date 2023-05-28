package checkers

import (
	"crypto/tls"
	"fmt"
	log "github.com/sirupsen/logrus"
	"host-cdn-checker/config"

	"net/http"
	"strings"
	"time"
)

type CertificateChecker struct {
}

func (c *CertificateChecker) Check(ip string) (isCDN bool, trust bool, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Timeout: 10 * time.Second, Transport: tr}

	seedUrl := fmt.Sprintf("https://%s", ip)
	resp, err := client.Get(seedUrl)
	defer resp.Body.Close()

	if err != nil {
		fmt.Errorf(seedUrl, " 请求失败")
		panic(err)
	}

	conf := config.Instance.Certificates

	//fmt.Println(resp.TLS.PeerCertificates[0])
	for _, certInfo := range resp.TLS.PeerCertificates {

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
