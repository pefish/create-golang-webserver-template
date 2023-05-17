
DEFAULT: build-cur

ifeq ($(GOPATH),)
  GOPATH = $(HOME)/go
endif

build-cur:
	GOPATH=$(GOPATH) go install github.com/pefish/go-build-tool/cmd/...@latest
	$(GOPATH)/bin/go-build-tool

install: build-cur
	sudo install -C ./build/bin/linux/create_golang_app_template /usr/local/bin/create_golang_app_template

install-service: install
	sudo mkdir -p /etc/systemd/system
	sudo install -C -m 0644 ./script/create_golang_app_template.service /etc/systemd/system/create_golang_app_template.service
	sudo systemctl daemon-reload
	@echo
	@echo "create_golang_app_template service installed."

