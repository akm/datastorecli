package main

import (
	"regexp"
	"strconv"

	"cloud.google.com/go/datastore"
	"github.com/akm/datastorecli"
)

var numberOnly = regexp.MustCompile(`\A\d+\z`)

func buildKey(args []string, encodedKey bool, encodedParent string) (*datastore.Key, error) {
	if encodedKey {
		key, err := datastorecli.DecodeKey(args[0])
		if err != nil {
			return nil, err
		}
		return key, nil
	} else {
		kind := args[0]

		parentKey, err := datastorecli.DecodeKey(encodedParent)
		if err != nil {
			return nil, err
		}

		if numberOnly.MatchString(args[1]) {
			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return nil, err
			}
			return datastore.IDKey(kind, id, parentKey), nil
		} else {
			return datastore.NameKey(kind, args[1], parentKey), nil
		}
	}
}
