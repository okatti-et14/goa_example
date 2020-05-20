// Code generated by goa v3.1.2, DO NOT EDIT.
//
// sample client
//
// Command:
// $ goa gen tmp-goa/design

package sample

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "sample" service client.
type Client struct {
	LoginEndpoint goa.Endpoint
}

// NewClient initializes a "sample" service client given the endpoints.
func NewClient(login goa.Endpoint) *Client {
	return &Client{
		LoginEndpoint: login,
	}
}

// Login calls the "login" endpoint of the "sample" service.
func (c *Client) Login(ctx context.Context, p *LoginPayload) (res *LoginResult, err error) {
	var ires interface{}
	ires, err = c.LoginEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LoginResult), nil
}
