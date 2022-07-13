package conf

import (
	"fmt"
	"testing"
	"tmpl-go-vercel/app/global"

	"github.com/spf13/viper"
)

func Test_Conf(t *testing.T) {
	configPath := "../conf/conf-dev.yaml"
	v := viper.New()
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(v.Get("mysql"))
	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println(global.Config)
	fmt.Println(global.Config.Name)
}
