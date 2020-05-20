package design

import . "goa.design/goa/v3/dsl"

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
			Attribute("result_tmp1", func() {
				Attribute("result_tmp2", Int)
			})
		})
		HTTP(func() {
			GET("/")
			Params(func() {
				Param("request_tmp")
			})
			Response(StatusOK)
		})
		Payload(func() {
			Attribute("request_tmp", Int)
			Required("request_tmp")
		})
	})
})
