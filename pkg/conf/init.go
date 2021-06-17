package conf

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"strings"
)
var (
	// Conf app global config
	Conf = &Config{}
)

// Init 初始化
func Init(configPath string) (*Config, error) {
	// 从本地读取环境变量
	godotenv.Load()
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	Conf, err = ParseConfig(cfgFile)
	return Conf, nil
}

// LoadConfig load config file from given path
func LoadConfig(confPath string) (*viper.Viper, error) {
	v := viper.New()
	if confPath != "" {
		v.SetConfigFile(confPath) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		v.AddConfigPath("config") // 如果没有指定配置文件，则解析默认的配置文件
		v.SetConfigName("config")
	}
	v.SetConfigType("yaml") // 设置配置文件格式为YAML

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// ParseConfig Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}