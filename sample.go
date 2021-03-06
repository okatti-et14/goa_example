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
	fake := "2020/12/10"
	res = &sample.LoginResult{
		ResultTmp3: &fake,
	}
	s.logger.Print("sample.login")
	return
}
