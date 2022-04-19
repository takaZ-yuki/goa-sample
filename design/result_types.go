package design

import . "goa.design/goa/v3/dsl"

// ユーザー情報
var User = Type("User", func() {
	Attribute("id", Int, "ユーザーID")
	Attribute("name", String, "ユーザー名", func() {
		MaxLength(20)
		Example("John Lennon")
	})
	Attribute("email", String, "Eメール", func() {
		MaxLength(255)
		Pattern(`/^[a-zA-Z0-9_.+-]+@([a-zA-Z0-9][a-zA-Z0-9-]*[a-zA-Z0-9]*\.)+[a-zA-Z]{2,}$/`)
		Example("xxxxxx@xxxx.xxx")
	})
})