package mysql

import (
	"log"
	"os"
	"time"
	"tmpl-go-vercel/app/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDb() (*gorm.DB, error) {
	var (
		dsn string
	)
	dev := global.Config.Dev
	logLv := logger.Warn
	colorful := false
	if dev {
		logLv = logger.Info
		colorful = true
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLv,
			Colorful:                  colorful,
			IgnoreRecordNotFoundError: false,
		},
	)

	if dev {
		dsn = "tcp(localhost:3309)/db?charset=utf8mb4&parseTime=True&loc=Local"
	} else {
		dsn = global.MysqlDsn
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err = sqlDb.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
