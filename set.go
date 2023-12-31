package itisadb

import (
	"context"
	"github.com/egorgasay/gost"
	api "github.com/egorgasay/itisadb-shared-proto/go"
)

var (
	setDefault int32 = 0
	setToAll   int32 = -1
)

func (c *Client) set(ctx context.Context, key, val string, opt SetOptions) (res gost.Result[int32]) {
	r, err := c.cl.Set(withAuth(ctx), &api.SetRequest{
		Key:   key,
		Value: val,
		Options: &api.SetRequest_Options{
			Server:   opt.Server,
			ReadOnly: opt.ReadOnly,
			Level:    api.Level(opt.Level),
			Unique:   opt.Unique,
		},
	})

	if err != nil {
		return res.Err(errFromGRPCError(err))
	}

	return res.Ok(r.SavedTo)
}

// SetOne sets the val for the key to gRPCis.
func (c *Client) SetOne(ctx context.Context, key, val string, opts ...SetOptions) (res gost.Result[int32]) {
	opt := SetOptions{}
	if len(opts) > 0 {
		opt = opts[0]
	}

	res = c.set(ctx, key, val, opt)
	if res.IsErr() {
		return res
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	server := res.Unwrap()
	c.keysAndServers[key] = server
	return res
}

// SetToAll sets the val for the key on all servers.
func (c *Client) SetToAll(ctx context.Context, key, val string, opts ...SetOptions) (res gost.Result[gost.Nothing]) {
	opt := SetOptions{
		Server: &setToAll,
	}

	if len(opts) > 0 {
		opt = opts[0]
	}

	r := c.set(ctx, key, val, opt)
	if r.IsErr() {
		return res.Err(r.Error())
	}

	return res.Ok(gost.Nothing{})
}

// SetMany sets a set of vals for gRPCis.
func (c *Client) SetMany(ctx context.Context, kv map[string]string, opts ...SetOptions) (res gost.Result[gost.Nothing]) {
	opt := SetOptions{}

	if len(opts) > 0 {
		opt = opts[0]
	}

	for key, val := range kv {
		r := c.set(ctx, key, val, opt)
		if r.IsErr() {
			return res.Err(r.Error())
		}
	}
	return res.Ok(gost.Nothing{})
}

// SetManyOpts gets a lot of vals from gRPCis with opts.
func (c *Client) SetManyOpts(ctx context.Context, keyValue map[string]Value) (res gost.Result[gost.Nothing]) {
	for key, val := range keyValue {
		r := c.set(ctx, key, val.Value, val.Options)
		if r.IsErr() {
			return res.Err(r.Error())
		}
	}
	return res.Ok(gost.Nothing{})
}
