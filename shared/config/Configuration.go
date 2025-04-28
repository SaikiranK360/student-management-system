package config

import (
	"os/user"

	"github.com/spf13/viper"
)

type Configuration struct {
	Environment Environment
	DB          DB
	Service     Service
}

var Config Configuration

func setConfig(v *viper.Viper) {
	Config.Environment = parseEnvironment(v.GetString("environment"))
	Config.DB = DB{
		Name:         v.GetString("db.name"),
		Username:     v.GetString("db.username"),
		PassWord:     v.GetString("db.password"),
		StudentDBUrl: v.GetString("db.studentDbUrl"),
	}
	Config.Service = Service{
		Broker: ServiceData{
			Url: v.GetString("service.broker.url"),
		},
		Student: ServiceData{
			Url: v.GetString("service.student.url"),
		},
	}
}

func LoadConfig() {
	currentUser, _ := user.Current()
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath(currentUser.HomeDir + "./test-config")

	// Reading the configuration info from the config.json file
	err := v.ReadInConfig()

	if err != nil {
		panic(err)
	}
	setConfig(v)
}
