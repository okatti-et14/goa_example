// Code generated by goa v3.1.2, DO NOT EDIT.
//
// sample HTTP client types
//
// Command:
// $ goa gen tmp-goa/design

package client

import (
	sample "tmp-goa/gen/sample"
)

// LoginResponseBody is the type of the "sample" service "login" endpoint HTTP
// response body.
type LoginResponseBody struct {
	ResultTmp1 *struct {
		ResultTmp2 *int `form:"result_tmp2,omitempty" json:"result_tmp2,omitempty" xml:"result_tmp2,omitempty"`
	} `form:"result_tmp1,omitempty" json:"result_tmp1,omitempty" xml:"result_tmp1,omitempty"`
}

// NewLoginResultOK builds a "sample" service "login" endpoint result from a
// HTTP "OK" response.
func NewLoginResultOK(body *LoginResponseBody) *sample.LoginResult {
	v := &sample.LoginResult{}
	if body.ResultTmp1 != nil {
		v.ResultTmp1 = &struct {
			ResultTmp2 *int
		}{
			ResultTmp2: body.ResultTmp1.ResultTmp2,
		}
	}

	return v
}