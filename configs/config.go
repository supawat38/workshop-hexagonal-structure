package configs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/nqd/flat"
	"github.com/spf13/viper"
)

type Config struct {
	App          `mapstructure:"app"`
	Postgres     `mapstructure:"postgres"`
	PostgresTest Postgres `mapstructure:"postgrestest"`
}

type App struct {
	Debug bool   `mapstructure:"debug"`
	Env   string `mapstructure:"env"`
	Port  string `mapstructure:"port"`
}

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DbName   string `mapstructure:"database"`
	SSLMode  bool   `mapstructure:"sslmode"`
}

var config Config

func InitViper(path, env string) {
	//Connect DB
	getLocalEnv()
}

func GetViper() *Config {
	return &config
}

func getLocalEnv() {

	secretData := map[string]interface{}{}

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		secretData[pair[0]] = os.Getenv(pair[0])
	}

	addEnv, err := flat.Unflatten(lower(secretData), &flat.Options{
		Delimiter: "_",
	})

	if err != nil {
		panic(err.Error())
	}

	jsonConfig, err := json.Marshal(addEnv)
	if err != nil {
		fmt.Println(err)
	}

	viper.SetConfigType("json")
	viper.AutomaticEnv()
	viper.ReadConfig(bytes.NewReader(jsonConfig))
	viper.Unmarshal(&config)

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println(err)
	}
}

func lower(v map[string]interface{}) map[string]interface{} {

	lv := make(map[string]interface{}, len(v))
	for mk, mv := range v {
		lv[strings.ToLower(mk)] = mv
	}
	return lv
}
