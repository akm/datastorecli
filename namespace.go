package datastorecli

import "cloud.google.com/go/datastore"

type Namespace string

func (s Namespace) PrepareKey(key *datastore.Key) *datastore.Key {
	if s != "" {
		key.Namespace = string(s)
	}
	return key
}

func (s Namespace) PrepareQuery(q *datastore.Query) *datastore.Query {
	if s != "" {
		return q.Namespace(string(s))
	}
	return q
}
