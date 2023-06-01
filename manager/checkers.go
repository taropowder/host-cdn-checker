package manager

import (
	"github.com/taropowder/host-cdn-checker/checkers"
)

type Checker interface {
	Check(ip string) (bool, bool, error)
}

func IsCDN(ip string) bool {
	checkersList := []Checker{
		&checkers.PortChecker{},
		&checkers.CertificateChecker{},
		&checkers.CDNCheckChecker{},
	}

	for _, checker := range checkersList {
		if isCdn, trust, err := checker.Check(ip); trust && err == nil {
			return isCdn
		}
	}

	return false
}
