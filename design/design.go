package design

import (
	. "goa.design/goa/v3/dsl"
	"goa.design/goa/v3/expr"
)

var _ = API("sample", func() {
	Title("sample API")
	Description("sampleAPI")

	HTTP(func() {
		Path("/")
	})

	Server("sample", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("sample", func() {
	Description("数値をそのまま返す")
	Method("login", func() {
		Result(func() {
			Attribute("XMLName", func() {
				Meta("struct:field:type", "xml.Name", "encoding/xml")
				Meta("struct:tag:xml", "cat")
			})
			Attribute("result_tmp1", func() {
				Attribute("result_tmp2", Int, func() {
					Meta("struct:tag:xml", "custcode_list,omitempty")
				})
			})
			Attribute("result_tmp3", String, func() {
				Format(expr.FormatDateTime)
			})
		})
		HTTP(func() {
			GET("/")
			Params(func() {
				Param("request_tmp")
			})
			Response(StatusOK, func() {
				ContentType("application/xml")
			})
		})
		Payload(func() {
			Attribute("request_tmp", Int)
			Required("request_tmp")
		})
	})
})
