package utils

import "github.com/asim/go-micro/v3/config"

// Mysql Config 結構
type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Port     string `json:"port"`
}

// 獲取 Mysql Consul 配置
func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	config.Get(path...).Scan(mysqlConfig)
	return mysqlConfig
}
