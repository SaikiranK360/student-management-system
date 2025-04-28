package config

//go:generate stringer -type=Environment
type Environment int

const (
	QA Environment = iota
	PROD
)

func parseEnvironment(env string) Environment {
	switch env {
	case "QA":
		return QA
	case "PROD":
		return PROD
	default:
		panic("Not able to parse the environment")
	}
}
