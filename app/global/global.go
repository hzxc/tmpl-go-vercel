package global

import (
	"crypto/rsa"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GlobalConfig struct {
	Name     string `mapstructure:"name" json:"name"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	Dev      bool   `mapstructure:"dev" json:"dev"`
}

var (
	Config    *GlobalConfig
	ZapLogger *zap.Logger
	PubKey    *rsa.PublicKey
	PrivKey   *rsa.PrivateKey
	MysqlDsn  string
	Db        *gorm.DB
)
