MAIN_PACKAGE_PATH=./cmd/datastorecli

VERSION=$(shell grep Version version.go | cut -f2 -d\")
TAG_NAME=v$(VERSION)

PACKAGES_ROOT_PATH=pkg
PACKAGES_PATH="$(PACKAGES_ROOT_PATH)/$(TAG_NAME)"

.PHONY: version
version:
	@echo $(VERSION)

.PHONY: build
build:
	go build -o /dev/null $(MAIN_PACKAGE_PATH)

GOX_PATH=$(GOPATH)/bin/gox
$(GOX_PATH):
	go get -u github.com/mitchellh/gox

.PHONY: build-packages
build-packages: $(GOX_PATH)
	gox \
		-arch=amd64 \
		-os=darwin \
		-os=linux \
		-os=windows \
		-output="$(PACKAGES_PATH)/{{.Dir}}-$(VERSION)-{{.OS}}_{{.Arch}}" \
		$(MAIN_PACKAGE_PATH)

GHR_PATH=$(GOPATH)/bin/ghr
$(GHR_PATH):
	go get -u github.com/tcnksm/ghr

.PHONY: release
release: $(GHR_PATH) build-packages
	ghr -draft $(TAG_NAME) $(PACKAGES_PATH)
