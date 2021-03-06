// Code generated by goa v3.1.2, DO NOT EDIT.
//
// sample HTTP server encoders and decoders
//
// Command:
// $ goa gen tmp-goa/design

package server

import (
	"context"
	"net/http"
	"strconv"
	sample "tmp-goa/gen/sample"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeLoginResponse returns an encoder for responses returned by the sample
// login endpoint.
func EncodeLoginResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*sample.LoginResult)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "application/xml")
		enc := encoder(ctx, w)
		body := NewLoginResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeLoginRequest returns a decoder for requests sent to the sample login
// endpoint.
func DecodeLoginRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			requestTmp int
			err        error
		)
		{
			requestTmpRaw := r.URL.Query().Get("request_tmp")
			if requestTmpRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("request_tmp", "query string"))
			}
			v, err2 := strconv.ParseInt(requestTmpRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("requestTmp", requestTmpRaw, "integer"))
			}
			requestTmp = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewLoginPayload(requestTmp)

		return payload, nil
	}
}
