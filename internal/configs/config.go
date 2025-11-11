package configs

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Environment   string `mapstructure:"ENVIRONMENT"`
	MaxCoin       int    `mapstructure:"MAX_COINS"`
	Port          string `mapstructure:"PORT"`
}


type BinanceConfig struct {
	Key    string `mapstructure:"BINANCE_KEY"`
	Secret string `mapstructure:"BINANCE_SECRET"`
}


type CoinConfig struct {
	Name string `mapstructure:"COIN_NAME"`
}


type DBConfig struct {
	DBName        string `mapstructure:"DB_NAME"`
	DBUsername    string `mapstructure:"DB_USERNAME"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
}



type Config struct {
	App AppConfig `mapstructure:"APP"`
	Binance BinanceConfig `mapstructure:"BINANCE"`
	Coin CoinConfig `mapstructure:"COIN"`
	DB DBConfig `mapstructure:"DB"`
}


// Load reads configuration from file or environment variables.
func Load(path string) (config Config, err error) { 
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {

		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("Error reading config file: %s. Relying on environment variables.", err)
		}
	}

	if err = viper.Unmarshal(&config); err != nil { 
		return config, fmt.Errorf("unmarshal config: %w", err)
	}

	return config, nil
}
