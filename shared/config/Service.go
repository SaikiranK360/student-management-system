package config

type ServiceData struct {
	Url  string
	Port string
	IP   string
}

type Service struct {
	Broker  ServiceData
	Student ServiceData
}

func (service Service) String() string {
	return ""
}

func (serviceData ServiceData) String() string {
	return ""
}
