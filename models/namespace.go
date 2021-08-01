package models

import (
	"context"
	"regexp"
	"strconv"

	"cloud.google.com/go/datastore"
)

type Namespace string

func (s Namespace) PrepareQuery(q *datastore.Query) *datastore.Query {
	if s != "" {
		return q.Namespace(string(s))
	}
	return q
}

func (s Namespace) BuildQuery(ctx context.Context, kind string, offset, limit int) *datastore.Query {
	q := datastore.NewQuery(kind).Limit(limit).Offset(offset)
	return s.PrepareQuery(q)
}

func (s Namespace) PrepareKey(key *datastore.Key) *datastore.Key {
	if s != "" {
		key.Namespace = string(s)
	}
	return key
}

var numberOnly = regexp.MustCompile(`\A\d+\z`)

func (s Namespace) BuildKey(args []string, encodedKey, incompleteKey bool, encodedParent string) (*datastore.Key, error) {
	if encodedKey {
		parentKey, err := DecodeKey(encodedParent)
		if err != nil {
			return nil, err
		}
		key := datastore.IncompleteKey(args[0], parentKey)
		return s.PrepareKey(key), nil
	} else if encodedKey {
		key, err := DecodeKey(args[0])
		if err != nil {
			return nil, err
		}
		return s.PrepareKey(key), nil
	} else {
		kind := args[0]

		parentKey, err := DecodeKey(encodedParent)
		if err != nil {
			return nil, err
		}

		var key *datastore.Key
		if numberOnly.MatchString(args[1]) {
			id, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return nil, err
			}
			key = datastore.IDKey(kind, id, parentKey)
		} else {
			key = datastore.NameKey(kind, args[1], parentKey)
		}
		return s.PrepareKey(key), nil
	}
}
