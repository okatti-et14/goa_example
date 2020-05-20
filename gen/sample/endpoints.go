// Code generated by goa v3.1.2, DO NOT EDIT.
//
// sample endpoints
//
// Command:
// $ goa gen tmp-goa/design

package sample

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "sample" service endpoints.
type Endpoints struct {
	Login goa.Endpoint
}

// NewEndpoints wraps the methods of the "sample" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Login: NewLoginEndpoint(s),
	}
}

// Use applies the given middleware to all the "sample" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Login = m(e.Login)
}

// NewLoginEndpoint returns an endpoint function that calls the method "login"
// of service "sample".
func NewLoginEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LoginPayload)
		return s.Login(ctx, p)
	}
}