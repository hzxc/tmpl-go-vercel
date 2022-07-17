package main

import (
	"fmt"
	"tmpl-go-vercel/app/global"
	model "tmpl-go-vercel/app/gorm/model"
	"tmpl-go-vercel/app/shared/mysql"

	"gorm.io/gorm"
)

func main() {
	var (
		db  *gorm.DB
		err error
	)
	global.Config = &global.GlobalConfig{Dev: true}

	if db, err = mysql.NewDb(); err != nil {
		goto ERR
	}

	if err = db.AutoMigrate(&model.Project{}); err != nil {
		goto ERR
	}

	if err = db.AutoMigrate(&model.Person{}); err != nil {
		goto ERR
	}

	return
ERR:
	fmt.Println(err.Error())
}
