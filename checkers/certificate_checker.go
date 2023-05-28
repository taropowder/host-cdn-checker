package checkers

import (
	"crypto/tls"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/taropowder/host-cdn-checker/config"
	"net/http"
	"strings"
	"time"
)

type CertificateChecker struct {
}

func (c *CertificateChecker) Check(ip string) (isCDN bool, trust bool, err error) {

	conf := config.Instance.Certificates

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Timeout: 10 * time.Second, Transport: tr}

	seedUrl := fmt.Sprintf("https://%s", ip)

	req, _ := http.NewRequest("GET", seedUrl, nil)
	// 增加 user-agent 请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {

		// 有的 CDN 无法用 InsecureSkipVerify 请求到, 正常请求拿报错中是否存在标志
		client = http.Client{Timeout: 10 * time.Second}
		req, _ := http.NewRequest("GET", seedUrl, nil)
		_, err := client.Do(req)
		if err != nil {
			for _, domain := range conf.BlackDomain {
				if strings.Contains(err.Error(), domain) {
					log.Info(fmt.Sprintf("certificate err  %s hit rule %s", err.Error(), domain))
					return true, true, nil
				}
			}

			for _, domain := range conf.WhiteDomain {
				if strings.Contains(err.Error(), domain) {
					return false, true, nil
				}
			}
		}

		return false, false, fmt.Errorf(seedUrl, " 请求失败", err)
	}

	defer resp.Body.Close()

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
