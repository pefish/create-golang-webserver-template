
DEFAULT: build-cur

ifeq ($(GOPATH),)
  GOPATH = $(HOME)/go
endif

build-cur:
	GOPATH=$(GOPATH) go install github.com/pefish/go-build-tool/cmd/...@latest
	go mod tidy
	$(GOPATH)/bin/go-build-tool

install: build-cur
	sudo install -C ./build/bin/linux/app-name /usr/local/bin/app-name

install-service: install
	sudo mkdir -p /etc/systemd/system
	sudo install -C -m 0644 ./script/app-name.service /etc/systemd/system/app-name.service
	sudo systemctl daemon-reload
	@echo
	@echo "app-name service installed."

