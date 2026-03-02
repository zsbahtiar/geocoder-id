VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS := -ldflags "-X github.com/zsbahtiar/geocoder-id.Version=$(VERSION)"

.PHONY: build build-all clean test

build:
	go build $(LDFLAGS) -o geocoder ./cmd/geocoder/

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o geocoder-darwin-amd64 ./cmd/geocoder/

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o geocoder-darwin-arm64 ./cmd/geocoder/

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o geocoder-linux-amd64 ./cmd/geocoder/

build-all: build-darwin-amd64 build-darwin-arm64 build-linux-amd64

clean:
	rm -f geocoder geocoder-*

test:
	go test -v ./...

run:
	go run $(LDFLAGS) ./cmd/geocoder/ geocode --coords="-6.2088 106.8456"

version:
	@echo $(VERSION)
