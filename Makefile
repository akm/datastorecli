MAIN_PACKAGE_PATH=./cmd/datastorecli

PACKAGES_PATH=pkg

VERSION=$(shell grep Version version.go | cut -f2 -d\")
TAG_NAME=v$(VERSION)

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: build
build:
	go build -o /dev/null $(MAIN_PACKAGE_PATH)

GOX_PATH=$(GOPATH)/bin/gox
$(GOX_PATH):
	go get -u github.com/mitchellh/gox

$(PACKAGES_PATH): build-packages

.PHONY: build-packages
build-packages: $(GOX_PATH)
	gox \
		-arch=amd64 \
		-os=darwin \
		-os=linux \
		-os=windows \
		-output="pkg/$(TAG_NAME)/{{.Dir}}-$(VERSION)-{{.OS}}_{{.Arch}}" \
		$(MAIN_PACKAGE_PATH)
