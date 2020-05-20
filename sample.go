package sampleapi

import (
	"context"
	"log"
	sample "tmp-goa/gen/sample"
)

// sample service example implementation.
// The example methods log the requests and return zero values.
type samplesrvc struct {
	logger *log.Logger
}

// NewSample returns the sample service implementation.
func NewSample(logger *log.Logger) sample.Service {
	return &samplesrvc{logger}
}

// Login implements login.
func (s *samplesrvc) Login(ctx context.Context, p *sample.LoginPayload) (res *sample.LoginResult, err error) {
	res = &sample.LoginResult{
		ResultTmp1: &struct{ ResultTmp2 *int }{1},
	}
	s.logger.Print("sample.login")
	return
}
