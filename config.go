package utils

import (
	"strconv"

	"github.com/asim/go-micro/plugins/config/source/consul/v3"
	"github.com/asim/go-micro/v3/config"
)

// 獲取 Consul 配置
func GetConsulConfig(host string, port int64, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		// 設定配置中心的地址
		// strconv 是我們轉換型別的函數
		// https://pkg.go.dev/strconv
		consul.WithAddress(host+":"+strconv.FormatInt(port, 10)),
		// 設定前綴，不設定 /micro/config
		consul.WithPrefix(prefix),
		// 是否移除前綴，設定true 表示可以不帶前綴直接獲取對應配置
		consul.StripPrefix(true),
	)

	// 初始化配置
	conf, err := config.NewConfig()
	if err != nil {
		return conf, err
	}
	// 加載配置
	err = conf.Load(consulSource)
	return conf, err
}
