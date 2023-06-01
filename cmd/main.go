package main

import (
	"fmt"
	"github.com/taropowder/host-cdn-checker/config"
	"github.com/taropowder/host-cdn-checker/manager"
)

func main() {
	// -p 命令行参数 ip 变量
	config.Instance = &config.Config{}
	config.Instance.Certificates.BlackDomain = []string{"cdn", "chinanetcenter"}
	fmt.Println(manager.IsCDN("103.254.188.41"))
}
