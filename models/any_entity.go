package models

import "cloud.google.com/go/datastore"

type AnyEntity map[string]interface{}

func (x AnyEntity) Load(props []datastore.Property) error {
	for _, prop := range props {
		x[prop.Name] = prop.Value
	}
	return nil
}

func (x AnyEntity) Save() ([]datastore.Property, error) {
	var props []datastore.Property
	for k, v := range x {
		props = append(props, datastore.Property{
			Name:  k,
			Value: v,
		})
	}
	return props, nil
}
