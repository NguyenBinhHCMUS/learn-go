package main

import "github.com/spf13/viper"

type Config struct {
	Database struct {
		MySQL struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
			Database string `mapstructure:"database"`
		} `mapstructure:"mysql"`
		Mongo struct {
			Host     string `mapstructure:"host"`
			Port     int    `mapstructure:"port"`
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
			Database string `mapstructure:"database"`
		} `mapstructure:"mongo"`
	} `mapstructure:"databases"`
	Secutity struct {
		Jwt struct {
			SecretKey     string `mapstructure:"secret_key"`
			ExpireMinutes int    `mapstructure:"expire_minutes"`
		} `mapstructure:"jwt"`
	} `mapstructure:"securities"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("development")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	host := viper.GetString("databases.mysql.host")
	port := viper.GetInt("databases.mysql.port")

	println("MySQL Host:", host)
	println("MySQL Port:", port)

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	println("MySQL Username:", config.Database.MySQL.Username)
	println("Mongo Database:", config.Database.Mongo.Database)
	println("JWT Secret Key:", config.Secutity.Jwt.SecretKey)
	println("JWT Expire Minutes:", config.Secutity.Jwt.ExpireMinutes)
}
