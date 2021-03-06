// Code generated by goa v3.1.2, DO NOT EDIT.
//
// sample HTTP client CLI support package
//
// Command:
// $ goa gen tmp-goa/design

package client

import (
	"fmt"
	"strconv"
	sample "tmp-goa/gen/sample"
)

// BuildLoginPayload builds the payload for the sample login endpoint from CLI
// flags.
func BuildLoginPayload(sampleLoginRequestTmp string) (*sample.LoginPayload, error) {
	var err error
	var requestTmp int
	{
		var v int64
		v, err = strconv.ParseInt(sampleLoginRequestTmp, 10, 64)
		requestTmp = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for requestTmp, must be INT")
		}
	}
	v := &sample.LoginPayload{}
	v.RequestTmp = requestTmp

	return v, nil
}
