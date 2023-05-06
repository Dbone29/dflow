package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Main    MainConfig
	Plugins map[string]viper.Viper
}

type MainConfig struct {
	Database DatabaseConfig
	Storage  StorageConfig
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type StorageConfig struct {
	Host            string
	AccessKeyID     string
	SecretAccessKey string
	BucketName      string
}

func LoadConfig() *Config {
	// Set configuration file options
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	// Attempt to read the configuration file
	_ = viper.ReadInConfig()

	// Set environment variable prefixes
	viper.SetEnvPrefix("DFLOW")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Unmarshal the configuration into corresponding structures
	mainConfig := MainConfig{}
	viper.UnmarshalKey("main", &mainConfig)

	pluginConfigs := make(map[string]viper.Viper)
	for pluginName := range viper.GetStringMap("plugins") {
		pluginViper := viper.Sub("plugins." + pluginName)
		pluginConfigs[pluginName] = *pluginViper
	}

	return &Config{
		Main:    mainConfig,
		Plugins: pluginConfigs,
	}
}
