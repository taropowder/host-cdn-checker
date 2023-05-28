package checkers

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type PortChecker struct {
}

func (p *PortChecker) Check(ip string) (isCDN bool, trust bool, err error) {
	// 扫描 ip 的 80 和 443 端口是否开放, 如果有一个没开放,则一定不是 CDN

	httpIsOpen := isOpen(ip, 80)
	httpsIsOpen := isOpen(ip, 443)
	if httpIsOpen && httpsIsOpen {
		return false, false, nil
	}

	return false, true, nil
}

// 查看端口号是否打开
func isOpen(ip string, port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), time.Second*time.Duration(5))
	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			fmt.Println("连接数超出系统限制！" + err.Error())
			os.Exit(1)
		}
		return false
	}
	_ = conn.Close()
	return true
}
