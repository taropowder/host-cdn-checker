package main

import (
	"fmt"
	"host-cdn-checker/config"
	"host-cdn-checker/manager"
)

func main() {
	// -p 命令行参数 ip 变量
	config.Instance = &config.Config{}
	config.Instance.Certificates.BlackDomain = []string{"cdn"}
	fmt.Println(manager.IsCDN("118.25.164.162"))
}