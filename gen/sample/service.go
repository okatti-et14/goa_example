// Code generated by goa v3.1.2, DO NOT EDIT.
//
// sample service
//
// Command:
// $ goa gen tmp-goa/design

package sample

import (
	"context"
	"encoding/xml"
)

// 数値をそのまま返す
type Service interface {
	// Login implements login.
	Login(context.Context, *LoginPayload) (res *LoginResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "sample"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"login"}

// LoginPayload is the payload type of the sample service login method.
type LoginPayload struct {
	RequestTmp int
}

// LoginResult is the result type of the sample service login method.
type LoginResult struct {
	XMLName    *xml.Name `xml:"cat"`
	ResultTmp1 *struct {
		ResultTmp2 *int `xml:"custcode_list,omitempty"`
	}
	ResultTmp3 *string
}
