package config

type Config struct {
	Certificates CertificatesConfig `json:"certificates"`
}

type CertificatesConfig struct {
	// 命中白名单则认为不是CDN
	WhiteDomain []string `json:"white_domain"`
	// 命中黑名单则认为是CDN
	BlackDomain []string `json:"black_domain"`
}

var Instance = &Config{}
