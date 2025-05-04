package config

type DB struct {
	Name         string
	Username     string
	PassWord     string
	StudentDBUrl string
	Charset      string
}

func (db DB) String() string {
	return ""
}
