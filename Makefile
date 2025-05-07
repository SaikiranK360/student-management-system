SERVICES := broker-service student-service
ENV_FILE := $(HOME)/.test-config/config.json

BRANCH := $(shell git name-rev --name-only HEAD)
COMMIT_ID := $(shell git rev-parse --short HEAD)
COMMIT_MESSAGE := $(shell git log -1 --pretty=%B)
BUILD_TIME := $(shell date)
ENVIRONMENT := $(shell cat $(ENV_FILE) | jq -r .environment)

build:
	cd proto && buf generate
	cd shared && go mod tidy
	@for service in $(SERVICES); do \
		( \
			cd service/$$service && \
			printf "<html>\n<body>\n<p>Environment: $(ENVIRONMENT)</p>\n<p>Branch: $(BRANCH)</p>\n<p>Commit id: $(COMMIT_ID)</p>\n<p>Commit message: $(COMMIT_MESSAGE)</p>\n<p>Build Time: $(BUILD_TIME)</p>\n</body>\n</html>" > version.html && \
			go mod tidy && \
			go build -o bin/ \
		); \
	done
