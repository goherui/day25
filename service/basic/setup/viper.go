package setup

import (
	"day25/service/basic/config"
	"fmt"

	"github.com/spf13/viper"
)

func InitViper() {
	var err error
	viper.SetConfigFile("../../../config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config.GlobalConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("配置加载成功")
}
