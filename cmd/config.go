package etl

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	MongoUri string `mapstructure:"MONGODB_URI"`
	DbName   string `mapstructure:"DB_NAME"`
	Region   string `mapstructure:"REGION"`
	Bucket   string `mapstructure:"BUCKET"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = viper.Unmarshal(&config)
	return
}
