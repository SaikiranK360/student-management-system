package config

type ServiceData struct {
	Url string
}

type Service struct {
	Broker  ServiceData
	Student ServiceData
}
