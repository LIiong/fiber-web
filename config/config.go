package config

import (
	"fiber-web/pkg/logger"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

var database *viper.Viper
var application *viper.Viper

//载入配置文件
func ConfigSetup(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Read config file[%s] fail: %s", path, err.Error()))
	}

	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		logger.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
	// 数据库初始化
	database = viper.Sub("db")
	if database == nil {
		panic("config not found settings.database")
	}
	DatabaseConfig = InitDatabase(database)
	// 服务器参数初始化
	application = viper.Sub("application")
	if application == nil {
		panic("config not found settings.database")
	}
	ApplicationConfig = InitApplication(application)
	// 日志配置
	logger.Init()
}

func SetConfig(configPath string, key string, value interface{}) {
	viper.AddConfigPath(configPath)
	viper.Set(key, value)
	_ = viper.WriteConfig()
}
