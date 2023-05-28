package checkers

import (
	"github.com/projectdiscovery/cdncheck"
	"net"
)

type CDNCheckChecker struct {
}

func (c *CDNCheckChecker) Check(ipString string) (isCDN bool, trust bool, err error) {
	client := cdncheck.New()
	ip := net.ParseIP(ipString)

	// checks if an IP is contained in the cdn denylist
	matched, _, err := client.CheckCDN(ip)
	if err != nil {
		return false, false, err
	}

	if matched {
		return true, true, nil
	}

	return false, false, nil
}
