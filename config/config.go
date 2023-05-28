package config

type Config struct {
	Certificates CertificatesConfig `json:"certificates"`
}

type CertificatesConfig struct {
	// 命中白名单则认为不是CDN
	WhiteDomain []string
	// 命中黑名单则认为是CDN
	BlackDomain []string
}

var Instance = &Config{}
