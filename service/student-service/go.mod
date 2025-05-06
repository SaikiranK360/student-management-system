module github.com/SaikiranK360/grpc-go-practice-2/service/student-service

go 1.24.2

require (
	github.com/SaikiranK360/grpc-go-practice-2/proto-gen v0.0.0
	github.com/SaikiranK360/grpc-go-practice-2/shared v0.0.0
	google.golang.org/grpc v1.72.0
	google.golang.org/protobuf v1.36.6
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.26.0
)

replace github.com/SaikiranK360/grpc-go-practice-2/shared => ../../shared

replace github.com/SaikiranK360/grpc-go-practice-2/proto-gen => ../../proto-gen

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/go-sql-driver/mysql v1.9.2 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/sagikazarmark/locafero v0.7.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.12.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/spf13/viper v1.20.1 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.9.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250505200425-f936aa4a68b2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
