// Code generated by goa v3.1.2, DO NOT EDIT.
//
// sample HTTP client types
//
// Command:
// $ goa gen tmp-goa/design

package client

import (
	"encoding/xml"
	sample "tmp-goa/gen/sample"

	goa "goa.design/goa/v3/pkg"
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

// NewLoginResultOK builds a "sample" service "login" endpoint result from a
// HTTP "OK" response.
func NewLoginResultOK(body *LoginResponseBody) *sample.LoginResult {
	v := &sample.LoginResult{
		XMLName:    body.XMLName,
		ResultTmp3: body.ResultTmp3,
	}
	if body.ResultTmp1 != nil {
		v.ResultTmp1 = &struct {
			ResultTmp2 *int `xml:"custcode_list,omitempty"`
		}{
			ResultTmp2: body.ResultTmp1.ResultTmp2,
		}
	}

	return v
}

// ValidateLoginResponseBody runs the validations defined on LoginResponseBody
func ValidateLoginResponseBody(body *LoginResponseBody) (err error) {
	if body.ResultTmp3 != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.result_tmp3", *body.ResultTmp3, goa.FormatDateTime))
	}
	return
}
