package models

import (
	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
)

func DecodeKey(encoded string) (*datastore.Key, error) {
	if encoded == "" {
		return nil, nil
	}
	if key, err := datastore.DecodeKey(encoded); err != nil {
		return nil, errors.Wrapf(err, "Failed to decode %s", encoded)
	} else {
		return key, nil
	}
}
