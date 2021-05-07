VERSION="0.0"
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_DATE=$(shell date --iso-8601)

WIFI_CONNECT_SOURCES=$(shell ls *.go)

ifeq ($(strip $(WIFI_ONBOARDING_VIEW_LOCATION)),)
WIFI_ONBOARDING_VIEW_LOCATION=./view/*
endif

ifeq ($(strip $(WIFI_ONBOARDING_STATIC_LOCATION)),)
WIFI_ONBOARDING_STATIC_LOCATION=./static
endif

ifeq ($(strip $(WIFI_ONBOARDING_DEFAULT_PORT)),)
WIFI_ONBOARDING_DEFAULT_PORT=":8080"
endif

DEPENDS=\
	github.com/gin-gonic/gin \
	github.com/PopcornComputer/gonnman

all: $(WIFI_CONNECT_SOURCES)
	@echo "Building Wifi-Connect"
	@echo "View Location: $(WIFI_ONBOARDING_VIEW_LOCATION)"
	@echo "Static Location: $(WIFI_ONBOARDING_STATIC_LOCATION)"
	@echo "Default Port: $(WIFI_ONBOARDING_DEFAULT_PORT)"
	@go build -o wifi-onboarding -ldflags="-s -w" -v \
	-ldflags "-X main.defaultPort=$(WIFI_ONBOARDING_DEFAULT_PORT) -X main.viewLocation=$(WIFI_ONBOARDING_VIEW_LOCATION) -X main.staticLocation=$(WIFI_ONBOARDING_STATIC_LOCATION)" \
	.
	@GOOS=linux GOARCH=arm go build -o build/linux_arm/wifi-onboarding -ldflags="-s -w" -v \
	-ldflags "-X main.defaultPort=$(WIFI_ONBOARDING_DEFAULT_PORT) -X main.viewLocation=$(WIFI_ONBOARDING_VIEW_LOCATION) -X main.staticLocation=$(WIFI_ONBOARDING_STATIC_LOCATION)" \
	.

clean:
	@echo "Cleaning"
	@rm -rf build/ wifi-onboarding

get:
	@echo "Downloading external dependencies"
	go get ${DEPENDS}
	@echo "Finished downloading external dependencies"
	@echo ${DEPENDS}
