module github.com/akm/datastorecli

go 1.16

require (
	cloud.google.com/go/datastore v1.5.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.2.1
)

exclude (
	github.com/hashicorp/go-version v1.3.0 // indirect
	github.com/mitchellh/gox v1.0.1 // indirect
)
