package configs

import (
	"log"

	"github.com/spf13/viper"
)

var Cfg Config

type Config struct {
	MongoUri string `mapstructure:"MONGODB_URI"`
	DbName   string `mapstructure:"DB_NAME"`
	Region   string `mapstructure:"REGION"`
	Bucket   string `mapstructure:"BUCKET"`
	Port     string `mapstructure:"PORT"`
	SentColl string `mapstructure:SENTCOLL`
}

func LoadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
}
