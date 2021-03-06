// Code generated by goa v3.1.2, DO NOT EDIT.
//
// sample HTTP server types
//
// Command:
// $ goa gen tmp-goa/design

package server

import (
	"encoding/xml"
	sample "tmp-goa/gen/sample"
)

// LoginResponseBody is the type of the "sample" service "login" endpoint HTTP
// response body.
type LoginResponseBody struct {
	XMLName    *xml.Name `xml:"cat"`
	ResultTmp1 *struct {
		ResultTmp2 *int `xml:"custcode_list,omitempty"`
	} `form:"result_tmp1,omitempty" json:"result_tmp1,omitempty" xml:"result_tmp1,omitempty"`
	ResultTmp3 *string `form:"result_tmp3,omitempty" json:"result_tmp3,omitempty" xml:"result_tmp3,omitempty"`
}

// NewLoginResponseBody builds the HTTP response body from the result of the
// "login" endpoint of the "sample" service.
func NewLoginResponseBody(res *sample.LoginResult) *LoginResponseBody {
	body := &LoginResponseBody{
		XMLName:    res.XMLName,
		ResultTmp3: res.ResultTmp3,
	}
	if res.ResultTmp1 != nil {
		body.ResultTmp1 = &struct {
			ResultTmp2 *int `xml:"custcode_list,omitempty"`
		}{
			ResultTmp2: res.ResultTmp1.ResultTmp2,
		}
	}
	return body
}

// NewLoginPayload builds a sample service login endpoint payload.
func NewLoginPayload(requestTmp int) *sample.LoginPayload {
	v := &sample.LoginPayload{}
	v.RequestTmp = requestTmp

	return v
}
