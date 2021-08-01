MAIN_PACKAGE_PATH=./cmd/datastorecli

.PHONY: build
build:
	go build -o /dev/null $(MAIN_PACKAGE_PATH)
