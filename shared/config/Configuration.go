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
		Charset:      v.GetString("db.charset"),
	}
	Config.Service = Service{
		Broker: ServiceData{
			Url:  v.GetString("service.broker.url"),
			Port: v.GetString("service.broker.port"),
			IP:   v.GetString("service.broker.ip"),
		},
		Student: ServiceData{
			Url:  v.GetString("service.student.url"),
			Port: v.GetString("service.student.port"),
		},
	}
}

func LoadConfig() {
	currentUser, _ := user.Current()
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("json")
	v.AddConfigPath(currentUser.HomeDir + "/.test-config")

	// Reading the configuration info from the config.json file
	err := v.ReadInConfig()

	if err != nil {
		panic(err)
	}
	setConfig(v)
}

/*
Overriding the default String() method
so that the config will not be printed anywhere
*/
func (c Configuration) String() string {
	return ""
}
