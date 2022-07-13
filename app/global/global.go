package global

import (
	"crypto/rsa"

	"go.uber.org/zap"
)

type GlobalConfig struct {
	Name     string      `mapstructure:"name" json:"name"`
	Username string      `mapstructure:"username" json:"username"`
	Password string      `mapstructure:"password" json:"password"`
	Dev      bool        `mapstructure:"dev" json:"dev"`
	Port     int         `mapstructure:"port" json:"port"`
	Mysql    MysqlConfig `mapstructure:"mysql" json:"mysql"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}

var (
	Config    *GlobalConfig
	PubKey    *rsa.PublicKey
	PrivKey   *rsa.PrivateKey
	ZapLogger *zap.Logger
)
