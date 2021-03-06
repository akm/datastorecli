package models

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
)

type Client struct {
	projectID string
	Namespace Namespace
}

func NewClient(projectID string, namespace string) *Client {
	return &Client{projectID: projectID, Namespace: Namespace(namespace)}
}

func (c *Client) dsClient(ctx context.Context) (*datastore.Client, error) {
	client, err := datastore.NewClient(ctx, c.projectID)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Put(ctx context.Context, key *datastore.Key, src interface{}) (*datastore.Key, error) {
	ds, err := c.dsClient(ctx)
	if err != nil {
		return nil, err
	}

	result, err := ds.Put(ctx, key, src)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to put %v with %s to %s", src, key.String(), key.Kind)
	}

	return result, nil
}

func (c *Client) Delete(ctx context.Context, key *datastore.Key) error {
	ds, err := c.dsClient(ctx)
	if err != nil {
		return err
	}
	if err := ds.Delete(ctx, key); err != nil {
		return errors.Wrapf(err, "failed to get by %s from %s", key.String(), key.Kind)
	}
	return nil
}

func (c *Client) Get(ctx context.Context, key *datastore.Key) (interface{}, error) {
	ds, err := c.dsClient(ctx)
	if err != nil {
		return nil, err
	}

	dst := AnyEntity{}
	if err := ds.Get(ctx, key, &dst); err != nil {
		return nil, errors.Wrapf(err, "failed to get by %s from %s", key.String(), key.Kind)
	}

	return dst, nil
}

func (c *Client) QueryKeys(ctx context.Context, kind string, offset, limit int) (*[]string, error) {
	ds, err := c.dsClient(ctx)
	if err != nil {
		return nil, err
	}
	q := c.Namespace.BuildQuery(ctx, kind, offset, limit)

	q = q.KeysOnly()
	keys, err := ds.GetAll(ctx, q, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to query from %s with keys only", kind)
	}
	r := make([]string, len(keys))
	for i, key := range keys {
		r[i] = key.Name
	}
	return &r, nil
}

func (c *Client) QueryData(ctx context.Context, kind string, offset, limit int) (*[]interface{}, error) {
	ds, err := c.dsClient(ctx)
	if err != nil {
		return nil, err
	}
	q := c.Namespace.BuildQuery(ctx, kind, offset, limit)

	var dst []AnyEntity
	if _, err := ds.GetAll(ctx, q, &dst); err != nil {
		return nil, errors.Wrapf(err, "failed to query from %s", kind)
	}
	r := make([]interface{}, len(dst))
	for i, x := range dst {
		r[i] = x
	}
	return &r, nil
}
