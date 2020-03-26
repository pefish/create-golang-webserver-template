
DEFAULT: all

all:
	rm -rf ./build/
	make build_linux

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -a -tags netgo -ldflags '-w -extldflags "-static"' -o build/main ./bin/main/
	@echo "Done building."

