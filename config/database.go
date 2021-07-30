package config

import "github.com/spf13/viper"

type Database struct {
	Host string
	Port string
	Database string
}

func InitDatabase(cfg *viper.Viper) *Database {
	return &Database{
		Port: cfg.GetString("port"),
		Host: cfg.GetString("host"),
		Database: cfg.GetString("database"),
	}
}

var DatabaseConfig = new(Database)
