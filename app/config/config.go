package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

type database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type app struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type config struct {
	Database database
	App      app
}

var cfg config

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func init() {
	viper.AddConfigPath(basepath)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config.yml")
	err := viper.MergeInConfig()
	if err != nil {
		panic(fmt.Errorf("Cannot load client config: %v", err))
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}
	viper.Unmarshal(&cfg)

	//fmt.Println("=============================")
	//fmt.Println(cfg)
	//fmt.Println("=============================")
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(env) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}

func GetConfig() *config {
	return &cfg
}
