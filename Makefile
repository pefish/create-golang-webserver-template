
DEFAULT: build-cur

ifeq ($(GOPATH),)
  GOPATH = $(HOME)/go
endif

build-cur:
	GOPATH=$(GOPATH) go install github.com/pefish/go-build-tool/cmd/...@latest
	$(GOPATH)/bin/go-build-tool

install: build-cur
	sudo install -C ./build/bin/linux/_app-name_ /usr/local/bin/_app-name_

install-service: install
	sudo mkdir -p /etc/systemd/system
	sudo install -C -m 0644 ./script/_app-name_.service /etc/systemd/system/_app-name_.service
	sudo systemctl daemon-reload
	@echo
	@echo "_app-name_ service installed."

