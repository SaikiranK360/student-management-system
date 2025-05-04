ENV_FILE := $(HOME)/.test-config/config.json

BRANCH := $(shell git name-rev --name-only HEAD)
COMMIT_ID := $(shell git rev-parse --short HEAD)
COMMIT_MESSAGE := $(shell git log -1 --pretty=%B)
BUILD_TIME := $(shell date)
ENVIRONMENT := $(shell cat $(ENV_FILE) | jq -r .environment)

build:
	cd ../../proto && buf generate
	printf "<html>\n<body>\n<p>Environment: $(ENVIRONMENT)</p>\n<p>Branch: $(BRANCH)</p>\n<p>Commit id: $(COMMIT_ID)</p>\n<p>Commit message: $(COMMIT_MESSAGE)</p>\n<p>Build Time: $(BUILD_TIME)</p>\n</body>\n</html>" > version.html && \
	go mod tidy
	go build -o ./bin/;

run: build
	./bin/student-service
