package datastorecli

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
)

type Client struct {
	projectID string
	namespace string
	kind      string
}

func NewClient(projectID string, namespace string, kind string) *Client {
	return &Client{projectID: projectID, namespace: namespace, kind: kind}
}

func (c *Client) dsClient(ctx context.Context) (*datastore.Client, error) {
	client, err := datastore.NewClient(ctx, c.projectID)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Get(ctx context.Context, key *datastore.Key) (interface{}, error) {
	ds, err := c.dsClient(ctx)
	if err != nil {
		return nil, err
	}

	dst := AnyEntity{}
	if err := ds.Get(ctx, key, &dst); err != nil {
		return nil, errors.Wrapf(err, "failed to get by %s from %s", key.String(), c.kind)
	}

	return dst, nil
}

func (c *Client) buildQuery(ctx context.Context, offset, limit int) *datastore.Query {
	q := datastore.NewQuery(c.kind).
		Limit(limit).
		Offset(offset)

	if c.namespace != "" {
		q = q.Namespace(c.namespace)
	}
	return q
}

func (c *Client) QueryKeys(ctx context.Context, offset, limit int) (*[]string, error) {
	ds, err := c.dsClient(ctx)
	if err != nil {
		return nil, err
	}
	q := c.buildQuery(ctx, offset, limit)

	q = q.KeysOnly()
	keys, err := ds.GetAll(ctx, q, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to query from %s with keys only", c.kind)
	}
	r := make([]string, len(keys))
	for i, key := range keys {
		r[i] = key.Name
	}
	return &r, nil
}

func (c *Client) QueryData(ctx context.Context, offset, limit int) (*[]interface{}, error) {
	ds, err := c.dsClient(ctx)
	if err != nil {
		return nil, err
	}
	q := c.buildQuery(ctx, offset, limit)

	var dst []AnyEntity
	if _, err := ds.GetAll(ctx, q, &dst); err != nil {
		return nil, errors.Wrapf(err, "failed to query from %s", c.kind)
	}
	r := make([]interface{}, len(dst))
	for i, x := range dst {
		r[i] = x
	}
	return &r, nil
}
