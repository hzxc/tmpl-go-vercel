package init

import (
	"encoding/json"
	"fmt"
	"os"
	"tmpl-go-vercel/app/global"
	"tmpl-go-vercel/app/shared/mysql"

	"go.uber.org/zap"

	"tmpl-go-vercel/app/utils/jwt"
)

func init() {
	var (
		zapLogger *zap.Logger
		err       error
	)
	fmt.Println("init zap logger")
	global.Config = &global.GlobalConfig{}
	json.Unmarshal([]byte(os.Getenv("VERCEL_ENV")), &global.Config)

	if global.Config.Dev {
		if zapLogger, err = zap.NewDevelopment(zap.AddCaller()); err != nil {
			zap.L().Fatal(err.Error())
		}
	} else {
		if zapLogger, err = zap.NewProduction(zap.AddCaller()); err != nil {
			zap.L().Fatal(err.Error())
		}
	}
	zap.ReplaceGlobals(zapLogger)
	global.ZapLogger = zapLogger

	fmt.Println("init jwt")
	PrivKeyEnv := os.Getenv("JWT_PRIVATE_KEY")
	global.PrivKey = jwt.ConvPrivKey(PrivKeyEnv)
	PubKeyEnv := os.Getenv("JWT_PUBLIC_KEY")
	global.PubKey = jwt.ConvPubKey(PubKeyEnv)

	fmt.Println("init mysql")
	if global.Config.Dev {
		global.MysqlDsn = os.Getenv("DATABASE_URL_DEV")
	} else {
		global.MysqlDsn = os.Getenv("DATABASE_URL")
	}

	if global.Db, err = mysql.NewDb(); err != nil {
		zap.L().Fatal(err.Error())
	}
}
